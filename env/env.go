package env


type Environment interface {
	GetArguments() []string
	AddArgument(arg string)
	SetArgument(index int, arg string)
	SetArguments(args []string)
}