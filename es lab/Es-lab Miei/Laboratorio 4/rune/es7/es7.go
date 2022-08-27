package main

import (
	"fmt"
)

func main() {
	var n string
	maiu, minus, vocal, cons := 0, 0, 0, 0

	fmt.Print("Inserisci la parola: ")
	fmt.Scan(&n)
	for _, i := range n {
		if èLetteraValida(i) {
			if èMaiuscola(i) {
				maiu++
			} else {
				minus++
			}
			if èVocale(i) {
				vocal++
			} else {
				cons++
			}
		}
	}
	fmt.Print("Maiuscole: ")
	fmt.Println(maiu)
	fmt.Print("Minuscole: ")
	fmt.Println(minus)

	fmt.Print("Vocali: ")
	fmt.Println(vocal)

	fmt.Print("Consonanti: ")
	fmt.Println(cons)

}
func test() {
	var n rune = 'I'
	fmt.Println(èVocale(n))

}

func èLetteraValida(l rune) bool {
	if l >= 'A' && l <= 'Z' || l >= 'a' && l <= 'z' {
		return true
	} else {
		return false
	}
}

func èMaiuscola(l rune) bool {
	if l >= 'A' && l <= 'Z' {
		return true

	} else {
		return false
	}
}

func èVocale(l rune) bool {
	if èMaiuscola(l) {
		l += 32
	}
	switch l {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
