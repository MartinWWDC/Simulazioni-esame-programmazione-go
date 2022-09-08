package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	mappa := make(map[string]int)
	for i := range str {
		for j := i; j <= len(str); j++ {
			sub := str[i:j]
			if isPalindomo(strings.Split(sub, "")) {
				mappa[sub]++
			}
		}
	}

	fmt.Println(mappa)
}

func isPalindomo(str []string) bool {
	var j int
	check := true
	if len(str) > 1 {
		for i := range str {
			j = len(str) - i - 1
			//fmt.Println(str[i], str[j])
			if str[i] != str[j] {
				check = false
			}
		}
	} else {
		check = false
	}
	return check
}
