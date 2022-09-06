package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text := LeggiTesto()
	s := Garibaldi(text)
	fmt.Println(s)
}

func LeggiTesto() (end string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			end += scanner.Text() + "\n"
		} else {
			break
		}
	}
	return
}
func TrasformaCarattere(c rune) rune {
	if strings.ContainsRune("aeioàáèéìíùúòó", unicode.ToLower(c)) {
		return 'u'
	} else {
		return c
	}
}

func Garibaldi(s string) (ns string) {
	ns = ""
	for _, g := range s {
		ns += string(TrasformaCarattere(g))
	}
	return
}
