package main

import (
	"fmt"
)

func main() {
	var v1, v2, tot float64
	var op rune
	flag := true

	fmt.Println("inserisci un'operazione aritmetica")
	fmt.Scanf("%f %c %f", &v1, &op, &v2)

	switch op {
	case '+':
		tot = (v1 + v2)
	case '-':
		tot = (v1 - v2)

	case '*':
		tot = (v1 * v2)

	case '/':
		tot = (v1 / v2)
	default:
		flag = false

	}

	if flag {
		fmt.Println("Risultato:", tot)
	} else {
		fmt.Println("Operatore non riconosciuto")
	}
}
