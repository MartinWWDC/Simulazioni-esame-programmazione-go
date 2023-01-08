package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	if isValid(s) {
		fmt.Println(getSubs(s))

	}
	//fmt.Println(isCresc(s))

}
func isValid(s string) bool {
	for _, k := range s {
		if !(k >= 'a' && k <= 'z') {
			return false
		}
	}
	return true

}
func getSubs(s string) map[string]int {
	mapp := make(map[string]int)
	for i := range s {
		for j := i + 2; j <= len(s); j++ {
			n := s[i:j]
			if isCresc(n) {
				mapp[n]++
			}
		}
	}
	return mapp
}
func isCresc(s string) bool {
	for j := 0; j < len(s)-1; j++ {
		if !(s[j] < s[j+1]) {
			return false
		}
	}
	return true
}
