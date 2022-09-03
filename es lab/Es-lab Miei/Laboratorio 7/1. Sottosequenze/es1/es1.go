package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(gen([]rune(os.Args[1])))
}

func gen(str []rune) []string {
	var arr []string
	for i := 0; i <= len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			if isPalindromo(str[i:j]) && len(str[i:j]) > 1 {
				arr = append(arr, string(str[i:j]))
			}
		}
	}
	return arr
}

func isPalindromo(str []rune) bool {
	unstr := ""
	for i := len(str) - 1; i >= 0; i-- {
		unstr += string(str[i])
	}
	return unstr == string(str)
}
