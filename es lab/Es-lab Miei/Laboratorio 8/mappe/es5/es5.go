package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	n := Occorrenze(LeggiTesto())
	fmt.Println("Occorrenze:")
	for k, v := range n {
		fmt.Printf("%c: %s\n", k, strings.Repeat("*", v))

	}
}
func LeggiTesto() (str string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			str += scanner.Text() + "\n"
		} else {
			break
		}
	}
	return
}

func Occorrenze(s string) map[rune]int {
	occorrenze := map[rune]int{}
	for _, l := range s {
		if unicode.IsLetter(l) {
			occorrenze[l]++
		}
	}
	return occorrenze
}
