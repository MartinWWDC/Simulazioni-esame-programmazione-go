/*
Scrivere un programma che legga da standard input un numero intero n e che
stampi a video un quadrato di lato n in cui:

una diagonale è rappresentata utilizzando il carattere o (lettera o);
l'area del quadrato al di sotto della diagonale è rappresentata con il carattere * (asterisco);
l'area del quadrato al di sopra della diagonale è rappresentata con il carattere + (più);
i caratteri di ogni riga del quadrato sono separati tra di loro da un carattere spazio.



input:
$ go run quadrato_con_diagonale.go
Inserisci un numero: 7
o + + + + + +
* o + + + + +
* * o + + + +
* * * o + + +
* * * * o + +
* * * * * o +
* * * * * * o


*/

package main

import "fmt"

func main() {
	var n, i, j int
	o := 0

	fmt.Scan(&n)

	for i = 0; i < n; i++ {

		for j = 0; j < n; j++ {

			if o != j {
				if j > o {
					fmt.Print("+")
				} else {
					fmt.Print("*")

				}
			} else {
				fmt.Print("o")
			}
			fmt.Print(" ")

		}
		o++
		fmt.Println("")

	}

}
