package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//k := LeggiNumero()
	//fmt.Println(k)
	fmt.Println((string(CifraCarattere('Z', 2))))
}
func LeggiNumero() (n int) {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	return
}

func LeggiTesto() (n string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n += scanner.Text()
	}
	return
}

func CifraCarattere(c rune, chiave int) rune {
	if (c >= 'A' || c <= 'Z') && (c >= 'a' || c <= 'z') {
		g := rune(int(c) + chiave)
		if g > 'Z' {
			fmt.Println("maiuscolo")
			g -= 'A'
		}
		if g > 'z' {
			fmt.Println("minuscolo")
			g -= 'a'
		}
		return g
	} else {
		return c
	}
}
