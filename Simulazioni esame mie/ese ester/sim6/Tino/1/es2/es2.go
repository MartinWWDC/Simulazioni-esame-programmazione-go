package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	g := 0
	for h := 0; h < n; h++ {
		for i := 0; i < 2*n+1; i++ {
			if i == n {
				if h == 0 {
					fmt.Print("+")

				} else {
					fmt.Print(" ")
				}
			} else {
				/*
					if i < n || (i >= n+g+1 && i+g-3 <= 2*n+1) {
						fmt.Print("o")
					} else {
						fmt.Print(" ")

					}*/
				if i >= n {
					if i-h > n {
						fmt.Print("o")
					} else {
						fmt.Print(" ")
					}
				} else {
					if n-i > h {
						fmt.Print("o")
					} else {
						fmt.Print(" ")
					}
				}

			}
		}
		g++
		fmt.Println("")
	}
}
