package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci il numero dei gradini: ")
	fmt.Scan(&n)
	StampaScala(n)

}

func StampaGradino(gradino int) {
	if gradino > 0 {
		for i := 1; i < 2*gradino; i++ {
			fmt.Print(" ")
		}
		fmt.Println("***")
		for i := 1; i < 2*gradino; i++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	} else {
		fmt.Println("Dimensione non sufficiente")
	}
}

func StampaScala(gradini int) {
	for i := gradini; i > 0; i-- {
		StampaGradino(i)
	}
}
