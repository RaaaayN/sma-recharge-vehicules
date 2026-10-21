package agt



//pour les agents on fait comme recommandé dans le cours, on utilise cette interface pour les faire évoluer dans l'environnement.
type Agent interface {
	Start()
	Percept(Environment)
	Deliberate()
	Act(Environment)
}


