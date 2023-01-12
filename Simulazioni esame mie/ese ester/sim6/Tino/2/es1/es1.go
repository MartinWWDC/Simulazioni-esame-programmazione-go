package main

import (
	"fmt"
	"unicode"
)

func main() {
	var pass string
	fmt.Scanln(&pass)
	fmt.Println(isAGoodPass(pass))
}

func isAGoodPass(pass string) bool {
	err := ""
	c := true
	if !(len(pass) >= 12) {
		c = false
		err += "- La pw deve avere una lunghezza minima di 12 caratteri\n"
	}

	nMin := 0
	nMaius := 0
	nNums := 0
	nChar := 0
	for _, h := range pass {
		if unicode.IsLetter(h) {
			if unicode.IsUpper(h) {
				nMaius++
			} else {
				nMin++
			}
		} else {
			if unicode.IsNumber(h) {
				nNums++
			} else {
				nChar++
			}
		}
	}
	if nMin < 2 {
		err += "- Almeno 2 caratteri della pw devono rappresentare delle lettere minuscole\n"
		c = false
	}
	if nMaius < 2 {
		err += "- Almeno 2 caratteri della pw devono rappresentare delle lettere maiuscole\n"
		c = false
	}
	if nNums < 3 {
		err += "- Almeno 3 caratteri della pw devono rappresentare delle cifre decimali\n"
		c = false
	}
	if nChar < 4 {
		err += "- Almeno 4 caratteri della pw non devono rappresentare lettere o cifre decimali\n"
		c = false
	}
	if !c {
		err = "La pw non è definita correttamente:\n" + err
	}
	fmt.Println(err)
	return c

}
