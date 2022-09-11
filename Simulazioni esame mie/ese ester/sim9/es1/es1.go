package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	g := checkpass(str)
	fmt.Println(g)
}

func checkpass(str string) bool {
	err := ""
	check := true
	maiusc, minusc, dec, div := 0, 0, 0, 0
	if len(str) < 12 {
		err += ("- La pw deve avere una lunghezza minima di 12 caratteri \n")
		check = false
	}
	for _, h := range str {
		if h >= '1' && h <= '9' {
			dec++
		} else {
			if h >= 'a' && h <= 'z' || h >= 'A' && h <= 'Z' {
				if unicode.IsUpper(h) {
					maiusc++
				} else {
					minusc++
				}
			} else {
				div++
			}
		}

	}
	if minusc < 2 {
		err += "- Almeno 2 caratteri della pw devono rappresentare delle lettere minuscole \n"
		check = false
	}
	if dec < 3 {
		err += "- Almeno 3 caratteri della pw devono rappresentare delle cifre decimali \n"
		check = false

	}
	if maiusc < 2 {
		err += "- Almeno 2 caratteri della pw devono rappresentare delle lettere maiuscole \n"
		check = false
	}
	if div < 4 {
		err += "- Almeno 4 caratteri della pw non devono rappresentare lettere o cifre decimali \n"
		check = false
	}
	if check {
		fmt.Println("La pw è ben definita!")
	} else {
		fmt.Println("La pw è non ben definita!")
		fmt.Println(err)

	}

	return check

}
