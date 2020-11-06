package helpers

import (
	"fmt"
	"strconv"
	"strings"

	model "github.com/sistemasnotaria4tlq/automatas-go/src/model"
)

func (af model.AutomataFinito) Validate(cadena string) (bool, string) {
	actual := af.init
	ActualState := strconv.Itoa(af.init)

	for _, item := range cadena {
		actual = af.mat[actual][strings.Index(af.texto, string(item))]
		ActualState = fmt.Sprintf("%s/%d", ActualState, actual)
	}

	return Include(af.valid, actual), ActualState
}

func Valid(af model.Automata, cadena string) {
	valid, ActualState := af.Validate(cadena)
	if valid {
		fmt.Println("ACCEPT!")
	} else {
		fmt.Println("REJECT!")
	}

	fmt.Println(ActualState)
}
func Include(s []int, e int) bool {
	for _, af := range s {
		if af == e {
			return true
		}
	}
	return false
}
