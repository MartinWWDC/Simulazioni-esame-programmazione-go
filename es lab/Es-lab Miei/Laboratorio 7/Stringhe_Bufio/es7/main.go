/*Hidden Number*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// leggo una singola riga di testo
	if scanner.Scan() {
		testo := scanner.Text()
		numero, err := NumeroNascosto(testo)
		// se la ricerca del numero nascosto è andata a buon fine stampo in output il numero nascosto
		if err == nil {
			fmt.Print("Doppio del numero nascosto: ", 2*numero, " (", numero, " * 2)\n")
		}
	}
}

func NumeroNascosto(testo string) (int, error) {
	/*
		il primo valore è un numero intero che rappresenta il numero nascosto all'interno del testo.
		Se il testo in input non contiene alcun numero il valore restituito deve essere 0;

		il secondo valore è l'eventuale errore restituito dalla funzione strconv.Atoi().
	*/
	var number string
	for _, carattere := range testo {
		if unicode.IsDigit(carattere) {
			number += string(carattere)
		}

	}
	return strconv.Atoi(number)
}
