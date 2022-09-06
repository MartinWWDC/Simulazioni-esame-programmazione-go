package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	n := LeggiTesto()
	nspazio := Spazia(n)
	fmt.Println(nspazio)
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func Spazia(s string) (ns string) {
	ns = ""
	for _, g := range s {
		if unicode.IsSpace(g) {
			ns += string(g)
		} else {
			ns += " " + string(g)
		}
	}
	return
}
