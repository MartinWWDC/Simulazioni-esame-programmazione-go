/*
legga da standard input un numero intero n
stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi
alternando fra loro colonne costituite da simboli * (asterisco) a colonne costituite da simboli + (più).

$ go run quadrato_colonne_alterne_1.go
Inserisci un numero: 5
* + * + *
* + * + *
* + * + *
* + * + *
* + * + *
*/

package main

import "fmt"

func main() {
	var n, i, j int
	fmt.Scanln(&n)
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if j%2 != 0 {

				fmt.Print("+")

			} else {

				fmt.Print("*")

			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
