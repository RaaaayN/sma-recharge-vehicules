/*
Pour avancer sur le GUI, faut créer le type snapshot qui va contenir toutes les infos qu'on veut afficher, et un type allocation qui va contenir les allocations de puissance pour chaque véhicule.


*/



package env 


// TODO Clarifier selon l'état de l'art, les ticks, les unités de temps, puissance etc ...
type TickSnapshot struct {
	T              int 
	PSite          float64
	TotalAllocated float64 
	Vehicles       []VehicleSnapshot
}


type VehicleSnapshot struct {
	ID           VehicleID 
	State        string
	Soc          float64
	AllocatedKW  float64 
	EnergyNeeded float64
}