package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
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

	cadena := ScanCadenas(scanner)

	texto := ScannerText(scanner)

	init, _ := strconv.Atoi(ScanCadenas(scanner))

	valid := scanFinalActualStates(scanner)

	mat := ScanMat(scanner)

	automata = &model.AFD{
		texto,
		init,
		valid,
		mat,
	}

	Valid(automata, cadena)
}
