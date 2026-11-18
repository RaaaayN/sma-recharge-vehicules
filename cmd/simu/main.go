
// C'est un script qui va lancer un serveur web qui va servir la GUI et des données factices pour tester la GUI.
// j'ai testé, et ça marche nickel, ça communique bien 

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"sma/env"
)


func fakeConfig() env.SimConfig {
	return env.SimConfig{
		Seed: 1, Strategy: "EqualShare", Ticks: 100, MinutesPerTick: 5,
		Vehicles: []env.VehicleConfig{
			{ID: "ev1", Arrival: 0, Departure: 60, EnergyNeeded: 30, MaxPower: 11},
			{ID: "ev2", Arrival: 5, Departure: 80, EnergyNeeded: 40, MaxPower: 22},
		},
	}
}

func fakeSnapshot(t int) env.TickSnapshot {
	const pSite = 20.0
	vs := []env.VehicleSnapshot{
		{ID: "ev1", State: "charging", Soc: min(1, 0.2+0.01*float64(t)), AllocatedKW: 10, EnergyNeeded: max(0, 30-0.8*float64(t))},
		{ID: "ev2", State: "charging", Soc: min(1, 0.1+0.008*float64(t)), AllocatedKW: 10, EnergyNeeded: max(0, 40-0.8*float64(t))},
	}
	total := 0.0
	for _, v := range vs {
		total += v.AllocatedKW
	}
	return env.TickSnapshot{T: t, PSite: pSite, TotalAllocated: total, Vehicles: vs}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fakeConfig())
	})

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		fl, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming non supporté", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")

		tick := time.NewTicker(time.Second)
		defer tick.Stop()
		for t := 0; ; t++ {
			select {
			case <-r.Context().Done():
				return
			case <-tick.C:
				b, _ := json.Marshal(fakeSnapshot(t))
				fmt.Fprintf(w, "data: %s\n\n", b)
				fl.Flush()
			}
		}
	})

	log.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}