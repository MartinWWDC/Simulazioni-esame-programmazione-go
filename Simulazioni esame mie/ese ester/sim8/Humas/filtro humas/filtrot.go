package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsLetter(rune(s[i])) && unicode.IsDigit(rune(s[i+1])) {
			fmt.Println(string(s[i]))
		}
	}
}
