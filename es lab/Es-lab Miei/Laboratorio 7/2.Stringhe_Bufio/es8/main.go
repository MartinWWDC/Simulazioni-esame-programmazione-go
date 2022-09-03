package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Inserisci un testo su più righe (termina con CTRL+D):")
	testo := LeggiTesto()
	fmt.Println(TraduciTestoInFarfallino(testo))
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}
func TraduciCarattereInFarfallino(c rune) string {
	/*
		una funzione TraduciCarattereInFarfallino(c rune) string che riceve in input un valore rune nel parametro c e restituisce un valore string che rappresenta la traduzione in linguaggio farfallino di c;
	*/
	if strings.ContainsRune("aeiou", c) {
		return string(c) + "f" + string(c)
	} else if strings.ContainsRune("AEIOU", c) {
		return string(c) + "F" + string(c)

	} else {
		return string(c)

	}

}

func TraduciTestoInFarfallino(s string) (testoUsci string) {
	for _, char := range s {
		testoUsci += TraduciCarattereInFarfallino(char)
	}
	return
}
