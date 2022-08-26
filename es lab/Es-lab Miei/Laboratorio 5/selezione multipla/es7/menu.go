/*
Scrivere un programma che permetta di ordinare al fast food con consegna a domicilio.

Il programma deve chiedere iterativamente da standard input il numero corrispondente al tipo di articolo (1 patatine,2 hamburger o 3 cocacola) e la quantità richiesta.

Con 0 viene terminata la lettura e viene stampato il costo dell'ordine.

I prezzi sono 2€ per le patatine 5€ per gli hamburger e 2€ per la cocacola.

In più ci sono 2€ di spesa per la consegna.

Si utilizzi il costrutto switch per selezionare l'articolo ed aggiornare il totale in base al numero inserito.

Input
$ go run menu.go
Cosa vuoi ordinare?
                1. patatine 2€
                2. hamburger 5€
                3. cocacola 2€
                0. termina
1
patatine? ottimo, quante?
3
Cosa vuoi ordinare?
                1. patatine 2€
                2. hamburger 5€
                3. cocacola 2€
                0. termina
2
hamburger? ottimo, quanti?
2
Cosa vuoi ordinare?
                1. patatine 2€
                2. hamburger 5€
                3. cocacola 2€
                0. termina
3
cocacola? ottimo, quante?
1
Cosa vuoi ordinare?
                1. patatine 2€
                2. hamburger 5€
                3. cocacola 2€
                0. termina
0
Sono 18 euro + 2 di consegna. Totale: 20
*/

package main

import "fmt"

func main() {
	var option int
	n := true
	total := 0
	for n != false {
		fmt.Println("vuoi ordinare?")
		fmt.Println("       1. patatine 2€")
		fmt.Println("       2. hamburger 5€")
		fmt.Println("       3. cocacola 2€")
		fmt.Println("       0. termina")
		fmt.Scanln(&option)
		switch option {
		case 1, 3:
			total += 2

		case 2:
			total += 5

		case 0:
			n = false

		default:
			fmt.Println("Non disponibile")
		}

	}

	fmt.Println("Sono ", total, " euro + 2 di consegna. Totale:", total+2)

}
