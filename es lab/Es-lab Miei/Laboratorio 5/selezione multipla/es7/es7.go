package main

import "fmt"

func main() {
	var n int
	conto := 0
	s := true
	for s {
		fmt.Println("Cosa vuoi ordinare?")
		fmt.Println("                1. patatine 2€")
		fmt.Println("                2. hamburger 5€")
		fmt.Println("                3. cocacola 2€")
		fmt.Println("                0. termina")
		fmt.Scan(&n)

		if n == 0 {
			s = false
		} else {
			conto += menu(n)

		}
	}
	fmt.Println("Sono", conto, "euro + 2 di consegna. Totale:", conto+2)
}

func menu(scelta int) int {
	switch scelta {
	case 1:
		fmt.Print("patatine?")
		return ripet(2)
	case 2:
		fmt.Print("hamburger?")

		return ripet(5)

	case 3:
		fmt.Print("cocacola?")
		return ripet(2)
	default:
		return 1

	}
}

func ripet(prezzo int) int {
	var n int
	fmt.Print("ottimo, quante?")
	fmt.Scan(&n)
	return prezzo * n

}
