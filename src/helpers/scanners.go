package helpers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ScanCadenas(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}

func ScannerText(scanner *bufio.Scanner) string {
	line := ScanCadenas(scanner)

	texto := ""
	for _, item := range strings.Split(line, ";") {
		if len(item) < 1 {
			continue
		}
		texto = fmt.Sprintf("%s%s", texto, item)
	}

	return texto
}

func scanFinalActualStates(scanner *bufio.Scanner) []int {
	line := ScanCadenas(scanner)
	valid := []int{}
	for _, item := range strings.Split(line, ";") {
		number, err := strconv.Atoi(item)
		if err != nil {
			continue
		}
		valid = append(valid, number)
	}

	return valid
}

func ScanMat(scanner *bufio.Scanner) [][]int {
	mat := [][]int{}
	for scanner.Scan() {
		array := []int{}
		for _, item := range strings.Split(scanner.Text(), ";") {
			number, err := strconv.Atoi(item)
			if err != nil {
				continue
			}
			array = append(array, number)
		}
		mat = append(mat, array)
	}

	return mat
}
