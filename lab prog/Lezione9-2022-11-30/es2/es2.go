package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in := os.Args[1]
	str := strings.Split(in, "")
	for i := 0; i <= len(str); i++ {
		for j := i; j <= len(str); j++ {
			sub := str[i:j]
			inv := []string{}
			for r := len(sub) - 1; r >= 0; r-- {
				inv = append(inv, string(sub[r]))
			}
			sub2 := sub
			//fmt.Println(inv)
			//fmt.Println(sub2)
			if Equal(inv, sub2) {
				//arr = append(arr, sub)
				fmt.Println(sub)
			}
		}
	}
}
func Equal(a, b []string) bool {
	if len(a) <= 1 && len(b) <= 1 || len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
