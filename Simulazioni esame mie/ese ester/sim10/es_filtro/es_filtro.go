package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]
	arrStr := strings.Split(str, "")
	counterlen := len(arrStr) - 1/2

	for i := len(arrStr) - 1; i >= 0; i-- {
		if i == counterlen/2 {
			fmt.Println(str)
		} else {
			for g := len(arrStr) / 2; g > 0; g-- {
				fmt.Print(" ")
			}
			fmt.Println((arrStr[i]))
		}

	}
}

/*
	counterlen := 0
	var invstr string
	for h, g := range str {
		counterlen = h
		invstr = string(str[g]) + invstr
	}
	fmt.Println(invstr)

	for i := 0; i < counterlen; i++ {
		if i == counterlen/2 {
			fmt.Println(str)
		} else {

			for g := 0; g < counterlen; g++ {
				if g == counterlen/2 {
					fmt.Print((invstr[i]))
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()

		}

	}
}
*/
