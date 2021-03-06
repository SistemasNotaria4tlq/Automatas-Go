package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/sistemasnotaria4tlq/automatas-go/src/helpers"
)

func main() {
	automata := Automata
	name := flag.String("file", "src/static/main.txt", "File name")
	flag.Parse()

	file, err := os.Open(*name)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cadena := helpers.ScanCadenas(scanner)

	texto := helpers.ScannerText(scanner)

	init, _ := strconv.Atoi(ScanCadenas(scanner))

	valid := helpers.scanFinalActualStates(scanner)

	mat := helpers.ScanMat(scanner)

	automata = &model.AFD{
		texto,
		init,
		valid,
		mat,
	}

	Valid(automata, cadena)
}
