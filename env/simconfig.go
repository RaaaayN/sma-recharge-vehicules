package env


type SimConfig struct {
	Seed           int64 
	Strategy       string 
	Ticks          int 
	MinutesPerTick int 
	Vehicles       []VehicleConfig
}

type VehicleConfig struct {
	ID           VehicleID 
	Arrival      int
	Departure    int 
	EnergyNeeded float64   
	MaxPower     float64   
}