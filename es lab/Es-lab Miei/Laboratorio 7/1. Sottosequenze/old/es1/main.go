/*
Scrivere un programma che legga da riga di comando una stringa composta di caratteri unicode e stampi a schermo tutte le sottostringhe
 palindrome (che siano uguali lette da destra e da sinistra) della stringa.

 Outuput
 $ go run esercizio_1.go sottotono
[otto tt tot oto ono]

$ go run esercizio_1.go banana
[ana anana nan ana]

$ go run esercizio_1.go ereggere
[ere ereggere regger egge gg ere]

$ go run esercizio_1.go Vaðlaheiðarvegavinnuverkfærageymsluskúrslyklakippuhringurinn
[nn pp nn]

$ go run esercizio_1.go donaudampfschifffahrtselektrizitätenhauptbetriebswerkbauunterbeamtengesellschaft
[ff fff ff ele izi tät uu ese ll]
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var arr []string
	stringa := os.Args[1]
	for i := 0; i < len(stringa); i++ {
		for j := i; j < len(stringa); j++ {
			if isPalindromo(stringa[i : j+1]) {

				//fmt.Println(stringa[i : j+1])
				arr = append(arr, stringa[i:j+1])
				//fmt.Println("palindroma")
			}

		}

	}
	fmt.Println(arr)
}

func isPalindromo(stringa string) bool {
	tes := false
	if len(stringa) > 1 {
		for i := 0; i < len(stringa); i++ {

			if stringa[i] == stringa[len(stringa)-i-1] {

				tes = true

			} else {
				tes = false
				i = len(stringa)

			}

		}
	}

	return tes
}
