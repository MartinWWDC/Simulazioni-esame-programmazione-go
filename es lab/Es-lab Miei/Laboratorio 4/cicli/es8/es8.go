package main

import "fmt"

func main() {
	var n, x int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			switch x {
			case 0:
				fmt.Print("* ")
			case 1:
				fmt.Print("+ ")
			case 2:
				fmt.Print("o ")

			}
		}
		fmt.Print(x)
		if x != 2 {
			x++
		} else {
			x = 0
		}
		fmt.Println()

	}
}
