package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(getSubs(os.Args[1]))
}
func isBalanced(s string) bool {
	oCounter := 0
	for _, i := range s {
		if i != '(' && i != ')' {
			return false
		}
		if i == '(' {
			oCounter++
		}
		if i == ')' {
			oCounter--
		}
		if oCounter < 0 {
			return false
		}
	}
	if oCounter != 0 {
		return false
	}
	return true
}
func getSubs(s string) map[string]int {
	mapp := make(map[string]int)
	for h := range s {
		for i := h + 1; i <= len(s); i++ {
			if isBalanced(s[h:i]) {
				mapp[s[h:i]]++
			}
		}
	}
	return mapp
}
