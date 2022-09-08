package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	n := LeggiNumero()
	k := LeggiTesto()
	fmt.Println(k)
	fmt.Println(CifraTesto(k, n))
}
func LeggiNumero() (n int) {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	return
}

func LeggiTesto() (n string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != " " {

			n += scanner.Text() + "\n"
		} else {
			fmt.Println("rompo")
			break
		}
	}
	return
}

func CifraCarattere(c rune, chiave int) rune {

	if (c >= 'A' || c <= 'Z') && (c >= 'a' || c <= 'z') && !unicode.IsSpace(c) {
		check := false
		if unicode.IsUpper(c) {
			check = true
		}
		g := rune(int(c) + chiave)

		if check && g > 'Z' {
			fmt.Println("maiuscolo")
			fmt.Println(g)
			g = 'A' + rune(chiave-1)
			fmt.Println(g)
		}
		if g > 'z' {
			fmt.Println("minuscolo")
			fmt.Println(g)
			g = 'a' + rune(chiave-1)
			fmt.Println(g)

		}
		return g

	} else {
		return c
	}
}

func CifraTesto(s string, chiave int) (final string) {
	final = ""
	for _, g := range s {
		fmt.Println(g)
		final += string(CifraCarattere(g, chiave))
	}
	return
}
