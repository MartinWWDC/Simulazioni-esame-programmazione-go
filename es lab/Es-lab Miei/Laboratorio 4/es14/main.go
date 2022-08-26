/*

Scrivere un programma che legga da standard input un intero n e

stampi a video un albero utilizzando il carattere * (asterisco) per rappresentarne la chioma ed il carattere + (più) per rappresentarne il tronco:

La chioma dell'albero deve essere alta n righe e, nel punto di larghezza massima, larga 2*n-1 colonne.
Il tronco dell'albero deve essere rappresentato con un quadrato di altezza e larghezza pari a 3 caratteri.
Il tronco dell'albero deve essere centrato rispetto alla chioma dell'albero.
Se n<=0, il programma stampa solo il tronco dell'albero.
7
      *
     ***
    *****
   *******
  *********
 ***********
*************
     +++
     +++
     +++
*/

package main

import "fmt"

func main() {
	var i, j, n, len int

	fmt.Scan(&n)

	if n == 0 || n == 1 {

		len = 3

	} else {

		len = (2 * n) - 1
		for i = 0; i < n; i++ {
			for j = 0; j < len; j++ {
				if j >= ((len/2)-i) && j <= ((len/2)+i) {
					fmt.Print("*")

				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println(" ")
		}
	}
	//tronco in extremiss
	for i = 0; i < 3; i++ {

		for j = 0; j < len; j++ {

			if j == (len/2)-1 || j == len/2 || j == len/2+1 {
				fmt.Print("+")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println("")
	}
}
