package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := strings.Split(os.Args[1], "")

	for h := range str {
		for g1, g := range str {
			if h == len(str)/2 {
				fmt.Print(string(g))

			} else {
				if g1 == len(str)/2 {
					fmt.Print(string(str[len(str)-h-1]))

				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println("")
	}
}
