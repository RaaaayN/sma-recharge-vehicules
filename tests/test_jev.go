package main

import (
	"fmt"
	"log"
	"os"

	"sma/env"
)

func main() {
	key := os.Getenv("TYPESAFE_API_KEY")
	if key == "" {
		log.Fatal("TYPESAFE_API_KEY manquant")
	}

	choice, confidence, err := env.NewJevClient(key).Classify(
		"Je n'ai plus que 20km d'autonomie, je ne sais pas si je vais pouvoir rentrer chez moi",
		map[string]string{
			"Recharger":  "Se rendre à une borne de recharge",
			"Eco": "Activer le mode éco pour économiser de l'énergie",
			"Conduire":  "Continuer à conduire jusqu'à la maison",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("choice=%s confidence=%.2f\n", choice, confidence)
}
