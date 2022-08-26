/*
Scrivere un programma che legga da standard input un intero n > 1
stampi, utilizzando il carattere *, il perimetro di due triangoli rettangoli con base e altezza di lunghezza n, posizionati come mostrato nell'Esempio d'esecuzione.





*/

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	// triangolo superiore

	var i, j int
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if i == 0 || j == n-1 || i == j {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	// triangolo inferiore
	for i = 0; i < n; i++ {
		for j = 0; j < n-1; j++ {
			fmt.Print(" ")
		}

		for j = 0; j < n; j++ {
			if i == n-1 || j == 0 || i == j {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
