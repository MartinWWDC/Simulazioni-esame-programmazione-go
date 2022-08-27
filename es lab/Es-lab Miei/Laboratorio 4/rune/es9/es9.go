package main

import "fmt"

func main() {
	var n string
	fmt.Print("Inserisci una stringa di testo: ")
	fmt.Scan(&n)
	for _, i := range n {
		fmt.Print(string(i))
		fmt.Print(" ")
	}
}
