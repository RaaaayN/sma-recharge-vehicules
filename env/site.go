//Définit les structures passives du lieu : Charger (avec sa puissance max) et ParkingSite (avec sa puissance max)

package env

// TODO vérifier après état de l'art fini
type ChargerKind int

const (
	AC ChargerKind = iota
	DC
)

// Charger
type Charger struct {
	ID       string
	Kind     ChargerKind
	MaxPower float64
	MinPower float64
	Occupant *int // pointeur vers l'ID du véhicule qui occupe la borne,passe à null si la borne est libre
}

// je sais pas si on le fait comme ça
type LoadProfile interface{ Power(t int) float64 }

type ConstantLoad struct{ Value float64 }

//TODO
//func Power(int)

// à vérifier aussi après état de l'art
type HourlyProfileLoad struct {
	Hourly [24]float64
	Noise  float64
}

//TODO
// func Power(t int)
