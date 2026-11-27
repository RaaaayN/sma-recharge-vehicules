package agt

import(
	"sma/config"
	"context"
	"fmt"
	"math"
)

//Définition de l'etat du vehicule (branché/débranché)

type VehicleState int

const(
	VehicleUnplugged VehicleState = iota //pas pluggé à une borne
	VehiclePlugged VehicleState // pluggé à une borne
)

type Vehicle struct{
	ID int

	//Batterie
	BatteryCapacity float64
	MaxPowerAcceptedKW float64
	CurrentSoC float64
	TargetSoC float64

	//arrivee et depart
	ArrivalTick int
	DepartTick int

	//Caractere du proprietaire du vehicule
	Alpha float64
	Beta float64
	Gamma float64

	//Etat spatial sur parking
	State VehicleState
	AssignedChargerID int

	//chanel pour recevoir 
	Cin chan Message

	//memoire pour lier les differentes etapes percept->delib->act
	lastMessage Message
	requestedPower float64
}


//lance routine autonome de l'agent vehcile
func (v *Vehicle) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // fin de la simu
			case msg := <-v.Cin:
				if msg.Type == ProposePrice {
					v.Percept(msg)
					v.Deliberate()
					v.Act()
				}
			}
		}
	}()
}

//Créer un nouveau vehicule et renvoyer son adresse memoire
func NewVehicle(id int, batCapKWh, maxPowAccKW, curSoc, targetSoc float64, arrTick, depTick int, alpha, beta, gamma float64) *Vehicle {
	return &Vehicle{
		ID: id,
		BatteryCapacity: batCapKWh,
		MaxPowerAcceptedKW: maxPowerAccKW,
		CurrentSoC: curSoC,
		TargetSoC: targetSoC,
		ArrivalTick: arrTick,
		DepartTick: depTick,
		Alpha: alpha,
		Beta: beta,
		Gamma: gamma,
		State: VehicleUnplugged,
		AssignedChargerID: -1,
		Cin: make(chan Message),
	}
}

//Percept -> Deliberate -> Act

//Percept
func (v *Vehicle) Percept(msg Message) {
	v.lastMessage = msg
}

//Deliberate 
func (v *Vehicle) Deliberate() {
	v.requestedPower = v.ComputeRequestedPower(v.lastMessage.Tick, v.lastMessage.Value)
}

//Act
func (v *Vehicle) Act() {
	if v.lastMessage.Reply != nil {
		v.lastMessage.Reply <- Message{
			From:  v.ID,
			Type:  PowerRequest,
			Tick:  v.lastMessage.Tick,
			Value: v.requestedPower,
		}
	}
}