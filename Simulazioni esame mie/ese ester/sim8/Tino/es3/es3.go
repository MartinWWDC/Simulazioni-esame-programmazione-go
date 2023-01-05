package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Studente struct {
	id      int
	nome    string
	cognome string
}

func main() {
	//maxC, _ := strconv.Atoi(os.Args[1])
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(parseStudenti(scanner.Text()))
	}

}
func creaStudente(id string) (studente Studente) {
	arr := strings.Split(id, ".")
	var idN int
	for j, h := range arr[1] {
		if unicode.IsNumber(h) {
			//fmt.Println(string(h))
			idN, _ = strconv.Atoi(arr[1][j:])

			arr[1] = arr[1][:j]
			break
		}
	}
	//fmt.Println(arr)

	studente = Studente{idN, arr[1], arr[0]}
	return
}
func parseStudenti(s string) (listaPrenotazione []Studente) {
	arr := strings.Split(s, " ")
	for _, h := range arr {
		listaPrenotazione = append(listaPrenotazione, creaStudente(h))
	}
	return
}
func printStudente(s Studente) string {
	return ""
}
