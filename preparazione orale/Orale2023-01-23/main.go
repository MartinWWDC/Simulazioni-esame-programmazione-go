package main

import (
	"fmt"
	"unicode"
)

/**
 * Prende una stringa e resituisce un int e conta le parole
 */
func main() {
	words := "ciao  come-u-ciao"
	fmt.Println(getWords(words))
}

func getWords(words string) int {
	g := 0
	for i := 0; i < len(words)-1; i++ {
		if i == 0 && unicode.IsLetter(rune(words[i])) {
			g++
		}
		if !unicode.IsLetter(rune(words[i])) && unicode.IsLetter(rune(words[i+1])) {
			g++
		}
	}
	return g
}
