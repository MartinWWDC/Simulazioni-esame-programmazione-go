package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		testo := scanner.Text()
		numero, err := NumeroNascosto(testo)
		//fmt.Println(numero)
		// se la ricerca del numero nascosto è andata a buon fine stampo in output il numero nascosto
		if err == nil {
			fmt.Print("Doppio del numero nascosto: ", 2*numero, " (", numero, " * 2)\n")
		}
	}
}

func NumeroNascosto(testo string) (int, error) {
	temp := ""
	var err error
	//var n int
	for _, g := range testo {
		if unicode.IsDigit(g) {
			if _, err = strconv.Atoi(string(g)); err == nil {
				temp += string(g)
			}
		}
	}
	nint, _ := strconv.Atoi(temp)

	return nint, err
}
