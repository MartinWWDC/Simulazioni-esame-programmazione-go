package main

import . "fmt"

func main() {
	s := "ƤƦŌĢɭăƁ"
	var str []rune
	for _, r := range s {
		str = append(str, r)
	}
	for i := len(str) - 1; i > 0; i-- {
		if i == len(str)/2 {
			Println(s)
		}
		for j := 0; j < len(str)/2; j++ {
			Print(" ")
		}
		Println(string(str[i]))

	}
}
