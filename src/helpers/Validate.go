package helpers

import (
	"fmt"
	"strconv"
	"strings"

	model "github.com/sistemasnotaria4tlq/automatas-go/src/model"
)

type AFD model.AutomataFinito

func (af *AFD) Validate(cadena string) (bool, string) {
	actual := af.Init
	ActualState := strconv.Itoa(af.Init)

	for _, item := range cadena {
		actual = af.Mat[actual][strings.Index(af.Texto, string(item))]
		ActualState = fmt.Sprintf("%s/%d", ActualState, actual)
	}

	return Include(af.Valid, actual), ActualState
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
