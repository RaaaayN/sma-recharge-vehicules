package main

import (
	"fmt"

	_ "sma/agt" // faux appels pour vérifier que les paquets agt et env compilent correctement
	_ "sma/env" // A MODIFIER : à supprimer quand on aura des appels réels à ces paquets
)

func main() {
	fmt.Println("Simulation de recharge de véhicules électriques : squelette OK")
}
