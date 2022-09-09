package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "password"
	isGood(s)
}

func isGood(str string) {
	fmt.Println(len(str))
	maiusc, minusc, nums, le := 0, 0, 0, 0
	check := true
	err := ""
	if len(str) != 8 {
		fmt.Println("- La pw deve avere una lunghezza di 8 caratteri")

	}
	for _, g := range str {

		if g > '0' && g < '9' {
			nums++

		} else {
			if !unicode.IsUpper(g) {
				minusc++
			}
			if unicode.IsUpper(g) {
				maiusc++
			}
		}

		if (g > '0' && g < '9') || (g >= 'a' && g <= 'z') || (g >= 'A' && g <= 'Z') {
			le++

		} else {
			if check {
				fmt.Println("- Tutti i caratteri della pw devono rappresentare lettere o cifre decimali")
			}
			check = false

		}

	}

	//fmt.Println(maiusc, minusc, nums, le)
	if minusc < 2 {
		fmt.Println("- Almeno 2 caratteri della pw devono rappresentare delle lettere minuscole")
	}
	if maiusc > 3 {
		fmt.Println("- Al massimo 3 caratteri della pw possono rappresentare delle lettere maiuscole")
	}
	if nums < 2 {
		//fmt.Println("num")
		fmt.Println("- Almeno 2 caratteri della pw devono rappresentare delle cifre decimali")
	}
	fmt.Println(err)
}
