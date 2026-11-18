package env

//TODO vérifier après état de l'art fini 
type ChargerKind int

const (
	AC ChargerKind = iota
	DC
)


// Charger 
type Charger struct {
	ID string
	Kind ChargerKind
	MaxPower float64
	MinPower float64
	Occupant *VehicleID
}


// je sais pas si on le fait comme ça
type LoadProfile interface {Power(t int) float64} 


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









type TariffSchedule interface{ GridPrice(t int) float64 }

type FlatTariff struct{ Price float64 }

//TODO
//func GridPrice(int)



//TODO prévoir un truc pour gérer les passages après minuit
//type TimeRange struct{} 

type TimeRange int

type PeakOffPeakTariff struct {
	OffPeakHours []TimeRange
	PeakPrice    float64
	OffPeakPrice float64
}


//TODO func GridPrice 













type Site struct {
	PConnection float64
	Chargers    []Charger
	Load        LoadProfile
	Tariff      TariffSchedule
}


//TODO func AvailablePower


//TODO func FreeCharger