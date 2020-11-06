package models

type Automata interface {
	Validate(string) (bool, string)
}
