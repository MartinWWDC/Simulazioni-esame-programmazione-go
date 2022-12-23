package main

import "fmt"

func main() {

	i := 1
	check := true
	for check {
		fmt.Print("Inserisci un numero: ")
		fmt.Scan(&i)
		check = Tebellina(i)

	}

}
func Tebellina(numero int) bool {
	if numero >= 1 && numero <= 9 {
		fmt.Print("Tabellina del ", numero, ": ")
		for g := 0; g <= 10; g++ {
			fmt.Print(g*numero, " ")
		}
		fmt.Println("")
		return true
	} else {
		fmt.Println("Programma Terminato.")
		return false
	}
}
