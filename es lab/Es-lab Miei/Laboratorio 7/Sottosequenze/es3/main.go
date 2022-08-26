/*

Scrivere un programma che:

legga da riga di comando una stringa s costituita da cifre decimali;
stampi a schermo tutte le sottosequenze (anche ripetute) della stringa s nelle quali le cifre sono in ordine crescente
(si considerino solamente sottosequenze di almeno 2 cifre).

Se la stringa s letta da riga di comando non è costituita solamente da cifre decimali,
il programma termina senza stampare nulla.


Oltre alla funzione main(), deve essere definita ed utilizzata almeno la funzione Sottostringhe(s string) []string,
che riceve in input un valore string nel parametro s,
e restituisce un valore []string che contiene tutte le sottosequenze desiderate

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var final []string
	for _, c := range os.Args[1] {
		if c < '0' || c > '9' {
			return
		}
	}
	input := os.Args[1]
	sub := Sottostringhe(input)
	//fmt.Println((sub))
	//fmt.Println(len(sub))

	for i := 0; i < len(sub); i++ {

		//fmt.Println("i:", i)

		//fmt.Println("sub:", string(sub[i]))

		check := false

		if len(sub[i]) > 1 {

			for j := 0; j < len(sub[i])-1; j++ {

				str := sub[i]

				//fmt.Println(string(str[j]), string(str[j+1]))

				if str[j] < str[j+1] && j != len(sub) {

					check = true

				} else {

					check = false

					j = len(sub)

				}
			}
			if check {
				//fmt.Println("fug")

				final = append(final, sub[i])

			}
		}
	}
	fmt.Println(final)

}

func Sottostringhe(s string) []string {
	var sub []string
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			sub = append(sub, s[i:j+1])
		}
	}
	return sub

}
