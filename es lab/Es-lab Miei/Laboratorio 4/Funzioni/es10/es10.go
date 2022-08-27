package main

import "fmt"

func main() {
	var tab int
	check := true
	for check {
		fmt.Print("Inserisci un numero: ")
		fmt.Scan(&tab)
		check = Tabellina(tab)
	}

}

func Tabellina(numero int) bool {
	if numero >= 1 && numero <= 9 {
		fmt.Print("Tabellina  del ")
		fmt.Print(numero)
		fmt.Print(": ")
		for i := 0; i <= 10; i++ {
			fmt.Print(numero * i)
			fmt.Print(" ")
		}
		fmt.Println()
		return true
	} else {
		fmt.Println("Numero non valido")
		return false
	}

}
