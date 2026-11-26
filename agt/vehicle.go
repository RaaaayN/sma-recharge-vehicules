package agt

// States
type VehicleState int

const (
	Waiting VehicleState = iota
	Charging
	Full
	Left
)
