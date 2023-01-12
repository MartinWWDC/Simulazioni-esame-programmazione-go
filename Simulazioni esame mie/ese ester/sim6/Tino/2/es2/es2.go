package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	r := 0
	for i := 0; i < n; i++ {

		for j := 0; j < 2*n+1; j++ {
			if j == n {
				if i == 0 {
					fmt.Print("+")

				} else {
					fmt.Print(" ")
				}
			} else {
				if j < n {
					if j < n-r {
						fmt.Print("0")
					} else {
						fmt.Print(" ")

					}
				} else {
					if j > n+r {
						fmt.Print("0")
					} else {
						fmt.Print(" ")

					}
				}
			}
		}
		r++
		fmt.Println("")
	}
}
