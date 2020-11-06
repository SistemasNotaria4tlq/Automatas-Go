package model

type Automata interface {
	Validate(string) (bool, string)
}
