package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < n; i++ {
		for g := 0; g < 2*n; g++ {

			if g == n {
				if i == 0 {
					fmt.Print("+")
				} else {
					fmt.Print(" ")
				}
			}
			if g >= n {
				if g-i >= n {
					fmt.Print(0)
				} else {

					fmt.Print(" ")
				}
			} else {
				if n-g > i {
					fmt.Print(0)
				} else {

					fmt.Print(" ")
				}
			}

		}
		fmt.Println()
	}
}
