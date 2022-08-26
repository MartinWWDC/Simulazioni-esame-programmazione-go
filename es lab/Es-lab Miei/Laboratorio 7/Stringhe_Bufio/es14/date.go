/*
Scrivere un programma che:

legga da standard input un testo su più righe;
termini la lettura quando viene inserita da standard input una riga vuota ( "" ).
Ogni riga di testo è una una stringa senza spazi che codifica una data in uno dei seguenti possibili formati:

aaaa/m/g
aaaa/m/gg
aaaa/mm/g
aaaa/mm/gg
g/m/aaaa
gg/m/aaaa
g/mm/aaaa
gg/mm/aaaa
Una volta terminata la fase di lettura, il programma deve stampare la codifica nel formato aaaa-mm-gg delle date lette. In particolare, le date devono essere stampate in ordine cronologico (dalla più antica alla più recente).

Oltre alla funzione main() deve essere definita ed utilizzata la funzione Formatta(data string) string, che data una data in input codificata in uno degli 8 formati descritti sopra, la restituisce formattata aaaa-mm-gg. La funzione deve utilizzare strings.Split per estrarre giorno, mese, e anno dalla stringa di input.

Suggerimento: Una volta codificate nel formato aaaa-mm-gg, le date possono essere ordinate cronologicamente utilizzando la funzione sort.Strings(a []string) del package sort.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	date := LeggiTesto()
	for _, i := range date {
		fmt.Println(Formatta(i))

	}
}
func LeggiTesto() (date []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() != "" {
			date = append(date, scanner.Text())
		} else {
			return
		}
	}
	return
}

func Formatta(data string) string {
	str := strings.Split(data, "/")
	var year, month, day string
	month = str[1]
	if len(str[0]) >= 4 {
		year = str[0]
		day = str[2]
	} else {
		year = str[2]
		day = str[0]
	}

	return year + "-" + month + "-" + day
}
