package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println(isGood(os.Args[1]))
}

func isGood(pass string) bool {
	err := ""

	nMinus := 0
	nMaiusc := 0
	nNums := 0
	nChar := 0
	for _, ch := range pass {
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			if unicode.IsLetter(ch) {

				if unicode.IsUpper(ch) {
					nMaiusc++
				} else {
					nMinus++
				}
				nChar++
			}
			if unicode.IsNumber(ch) {
				nNums++
			}
		} else {
			err += ("- Tutti i caratteri della pw devono rappresentare lettere o cifre decimali")
		}

	}
	c := true
	fmt.Println(nChar)
	if len(pass) != 8 {
		err += ("- La pw deve avere una lunghezza di 8 caratteri \n")
		c = false
	}
	if nMinus < 2 {
		err += ("- Almeno 2 caratteri della pw devono rappresentare delle lettere minuscole		\n")
		c = false
	}
	if nMaiusc > 3 {
		err += ("- Almeno 3 caratteri della pw devono rappresentare delle lettere maiuscole		\n")
		c = false
	}
	if nNums < 2 {
		err += ("- Almeno 2 caratteri della pw devono rappresentare delle cifre decimali		\n")
		c = false
	}

	if !c {
		fmt.Println("La pw non è definita correttamente:			")
		fmt.Println(err)
	} else {
		fmt.Println("La Passowrd è ben Definita")
	}
	return c
}
