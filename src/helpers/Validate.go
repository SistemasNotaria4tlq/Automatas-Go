package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func (a models.AFD) Validate(cadena string) (bool, string) {
	actual := a.init
	ActualState := strconv.Itoa(a.init)

	for _, item := range cadena {
		actual = a.mat[actual][strings.Index(a.texto, string(item))]
		ActualState = fmt.Sprintf("%s/%d", ActualState, actual)
	}

	return Include(a.valid, actual), ActualState
}

func Valid(a Automata, cadena string) {
	valid, ActualState := a.Validate(cadena)
	if valid {
		fmt.Println("ACCEPT!")
	} else {
		fmt.Println("REJECT!")
	}

	fmt.Println(ActualState)
}
func Include(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
