//interface des agents
//définition des messages échangés entre les agents pour la négo

package agt

import (
	"context"
)

type Agent interface {
	Start(ctx context.Context)
}

// Ici on définit le type de messages qui vont être échangés entre le manager et les vehicles.

type MessageType int

const (
	ProposePrice MessageType = iota //ParkingManager --> Vehicle
	PowerRequest                    // Vehicle --> ParkingManager
)

type Message struct {
	From  int //ID de l'émetteur
	Type  MessageType
	Tick  int //Pour que le vehicule sache quelle heure il est au moment de calculer le temps qu'il lui reste
	Value float64
	Reply chan<- Message //En ecriture seulement car une voiture ne doit pas pouvoir lire les reponses des autres voitures
}
