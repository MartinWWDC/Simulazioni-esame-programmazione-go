package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print("+")
			}

		}
		fmt.Println()
	}
}
