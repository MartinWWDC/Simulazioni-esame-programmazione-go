package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	txt := LeggiTesto()
	fmt.Println(TraduciTestoInFarfallino(txt))
}

func LeggiTesto() (txt string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		} else {
			txt += scanner.Text() + "\n"
		}
	}
	return
}
func TraduciCarattereInFarfallino(c rune) (carattereTrasformato string) {
	carattereTrasformato = string(c)

	if strings.ContainsRune("aeiouèéòóùúàáíì", c) {
		carattereTrasformato += "f" + carattereTrasformato
	}
	if strings.ContainsRune("AEIOUÀÁÈÉÌÍÓÒÚÙ", c) {
		carattereTrasformato += "F" + carattereTrasformato
	}
	return
}
func TraduciTestoInFarfallino(s string) (traduzione string) {
	for _, g := range s {
		traduzione += TraduciCarattereInFarfallino(g)
	}
	return traduzione
}
