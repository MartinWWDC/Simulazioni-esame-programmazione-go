package main

import "fmt"

func main() {
	var tab int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&tab)
	Tabellina(tab)

}

func Tabellina(numero int) {
	if numero >= 1 && numero <= 9 {
		fmt.Print("Tabellina  del ")
		fmt.Print(numero)
		fmt.Print(": ")
		for i := 0; i <= 10; i++ {
			fmt.Print(numero * i)
			fmt.Print(" ")
		}
	}

}
