package env

type VehicleID string


type EVView struct {
	ID VehicleID
	Arrival int
	Departure int
	EnergyNeeded float64
	Soc float64
	MaxPower float64
	MinPower float64
	Bid float64
}



//on part sur ça pour le moment
type Allocation map[VehicleID]float64

type PowerStrategy interface {
	Name() string
	Allocate(t int, pSite float64, evs []EVView) Allocation
}