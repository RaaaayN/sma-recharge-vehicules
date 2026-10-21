package agt


type Agent interface {
	Start()
	Percept(Environment)
	Deliberate()
	Act(Environment)
}
