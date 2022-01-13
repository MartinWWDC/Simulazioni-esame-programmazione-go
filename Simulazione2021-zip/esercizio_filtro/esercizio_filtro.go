/*
Scrivere un programma che legga da riga di comando un numero naturale
positivo e, come mostrato nell’Esempio di esecuzione, stampi a video ogni
cifra del numero inferiore a quella che la segue.
Si assuma un ordinamento delle cifre da sinistra verso destra. Ad esempio, dato
il numero 496:
•4 è la prima cifra del numero; la cifra 4 è seguita dalla cifra 9;
•9 è la seconda cifra del numero; la cifra 9 è seguita dalla cifra 6;
•6 è la terza e ultima cifra del numero; la cifra 6 non è seguita da nessun
altra cifra.
Si assuma inoltre che il valore letto da riga di comando sia nel formato corretto.
Si noti che l’ultima cifra del numero letto non viene mai stampata.

*/
package main

import (
	"fmt"
	"os"
)

func main() {
	num := os.Args[1]
	for i := 0; i < len(num)-1; i++ {
		//fmt.Println(int(num[i]), " < ", int(num[i+1]))
		if int(num[i]) < int(num[i+1]) {
			fmt.Print(string(num[i]))
		}
	}
}
