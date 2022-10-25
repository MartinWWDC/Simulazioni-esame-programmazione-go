package main

import "fmt"

func main() {
	var n int
	//fmt.Scan(&n)
	n = 11
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || j == n-1 || i == n-1 || j == n/2 || i == n/2 || i == j || i == n-j-1 {
				fmt.Print("*")

			} else {
				fmt.Print(" ")

			}

		}
		fmt.Println()
	}

}
