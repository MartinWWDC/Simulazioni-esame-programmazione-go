/*

Scrivere un programma che:
• legga da standard input una stringa s. Si assuma che la stringa non
contenga caratteri di spaziatura.

• verifichi che la stringa sia composta di sole lettere minuscole nello standard
di caratteri ASCII. Qualora non fosse così il programma deve terminare
senza stampare nulla.

• stampi a schermo, tutte le sottosequenze della stringa s nelle quali le lettere
siano in ordine crescente (si considerino solamente sottosequenze di almeno
2 lettere).


Come mostrato nell’Esempio d’esecuzione, ciascuna sottosequenza deve essere
stampata un’unica volta, riportando il numero di volte in cui la sottosequenza
appare in s.


Oltre alla funzione main(), deve essere definita ed utilizzata almeno la funzione
Sottostringhe(s string) map[string]int, che riceve in input un valore string nel parametro s,
e restituisce un valore map[string]int in cui, per
ogni sottosequenza di s di almeno 2 lettere, sia memorizzato il numero di volte
in cui essa appare in s.

*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text := os.Args[1]
	textLow := strings.ToLower(text)
	if text == textLow {
		fmt.Println("no upper")
		fmt.Println(Sottostringhe(text))
	}
}
func Sottostringhe(s string) map[string]int {
	mapAssigned := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] > s[j-1] {
				mapAssigned[s[i:j+1]]++
			} else {
				break
			}
		}
	}
	return mapAssigned
}
