package main

import (
	"fmt"
	"log"
	"os"

	"sma/env"
)

func main() {
	key := os.Getenv("GROQ_API_KEY")
	if key == "" {
		log.Fatal("GROQ_API_KEY manquant")
	}

	out, err := env.NewGroqClient(key).ChatCompletion("Salut !")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
