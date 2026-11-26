package env

type VehicleID string

type EVView struct {
	ID           VehicleID
	Arrival      int
	Departure    int
	EnergyNeeded float64
	Soc          float64
	MaxPower     float64
	MinPower     float64
	Bid          float64
}

//TODO faut voir comment on gère les allocations, on fait une map vehicle vers float ??
type Allocation float64

type PowerStrategy interface {
	Name() string
	Allocate(t int, pSite float64, evs []EVView) Allocation
}
