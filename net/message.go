//dans ce dossier, on va mettre tous les trucs qui concernent la communication entre les agents et l'environnement, donc les messages qu'ils s'échangent.


package net

import "sma/env"

type Tick struct{ 
	T int 
	PSite float64 
}



type Grant struct{ 
	KW float64 
}


type Request struct {
	View  env.EVView
	Reply chan<- Grant // à créer avec un buffer de 1
}