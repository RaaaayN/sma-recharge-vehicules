package env

type Persona string

type VehicleProfile struct {
	Arrival          int
	PlannedDeparture int
	BatteryCapacity  float64 // à clarfiier si on fait en kWh ou en kW
	SoCArrival       float64
	SoCTarget        float64
	MaxPower         float64 // idem
	Budget           float64
	Persona          Persona
}

type ScenarioSource interface {
	Profiles(seed int64) []VehicleProfile
}