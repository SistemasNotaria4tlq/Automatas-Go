package models

type AutomataFinito struct {
	Texto string
	Init  int
	Valid []int
	Mat   [][]int
}
