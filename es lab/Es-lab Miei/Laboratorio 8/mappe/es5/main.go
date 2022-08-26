/*
Scrivere un programma che:

legga da standard input un testo su più righe (alcune delle quali possono essere delle righe vuote (""));
termini la lettura quando, premendo la combinazione di tasti Ctrl+D, viene inserito da standard input l'indicatore End-Of-File (EOF);
come mostrato nell'Esempio di esecuzione, stampi un istogramma a barre orizzontali per rappresentare il numero di occorrenze di ogni lettera presente nel testo letto:
una lettera è un carattere il cui codice Unicode, se passato come argomento alla funzione func IsLetter(r rune) bool del package unicode, fa restituire true alla funzione;
le lettere minuscole sono da considerarsi diverse dalle lettere maiuscole;
ogni barra viene rappresentata utilizzando il carattere asterisco (*); se il numero di occorrenze della lettera e è per esempio 9, la barra corrispondente sarà formata da 9 caratteri *.
Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione LeggiTesto() string che legge da standard input un testo su più righe (alcune delle quali possono essere delle righe vuote ("")) e terminato dall'indicatore EOF, restituendo un valore string in cui è memorizzato il testo letto;
una funzione Occorrenze(s string) map[rune]int che riceve in input un valore string nel parametro s e restituisce un valore map[rune]int in cui, per ogni lettera presente in s, è memorizzato il numero di occorrenze della lettera in s.


*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	s := LeggiTesto()
	mappa := (Occorrenze(s))

	genGraph(mappa)

}

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	testo := ""
	for scanner.Scan() {
		testo += scanner.Text()
	}
	return testo
}

func Occorrenze(s string) map[rune]int {
	data := make(map[rune]int)
	/*
	   una funzione Occorrenze(s string) map[rune]int che riceve in input un valore string nel parametro s e restituisce un valore map[rune]int
	   in cui, per ogni lettera presente in s, è memorizzato il numero di occorrenze della lettera in s
	*/
	for _, letter := range s {
		if unicode.IsLetter(letter) {
			if _, ok := data[letter]; ok {
				data[letter] += 1

			} else {
				data[letter] = 1

			}
		}

	}
	return data
}
func genGraph(mappa map[rune]int) {
	for item := range mappa {
		fmt.Print(string(item), ":")
		for i := 0; i < mappa[item]; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
