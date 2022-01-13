package main

import "fmt"

func main() {
	fmt.Println(romano2decimale("MCMLXXXIII"))
}
func change(n string) []int {
	var arr []int
	for char := range n {
		switch {
		case string(n[char]) == "I":
			arr = append(arr, 1)
		case string(n[char]) == "V":
			arr = append(arr, 5)
		case string(n[char]) == "X":
			arr = append(arr, 10)
		case string(n[char]) == "L":
			arr = append(arr, 50)
		case string(n[char]) == "C":
			arr = append(arr, 100)
		case string(n[char]) == "D":
			arr = append(arr, 500)
		case string(n[char]) == "M":
			arr = append(arr, 1000)

		}
	}
	return arr
}

func romano2decimale(n string) int {
	/*
	   Il simbolo I può comparire alla sinistra dei simboli V e X, ma una sola volta. Ad esempio IX equivale a 9. IIX non è un
	   numero romano ben formattato, in quanto il numero romano corrispondente a 8 è VIII. Il simbolo X può comparire alla
	   sinistra dei simboli L e C, ma una sola volta. Ad esempio, il numero 99 corrisponde a XCIX: la somma di XC (90) e IX
	   (9). Il numero romano IC non è ben formattato. Infine, il simbolo C può comparire alla sinistra dei simboli D e M, ma
	   una sola volta. Ad esempio, il numero 900 corrisponde a CM. I simboli V, L e D non possono mai essere posizionati
	   alla sinistra di simboli con un valore più grande (VC ad esempio non è ben formattato, in quanto il numero 95
	   corrisponde a XCV: XC + V).
	*/
	toInt := change(n)
	sum := 0
	for i := 1; i < len(toInt); i++ {
		fmt.Println(toInt[i-1], ",", toInt[i])
		if toInt[i-1] >= toInt[i] {
			fmt.Println("add")

			sum += toInt[i-1]
		} else {
			sum += (toInt[i] - toInt[i-1])
			i++
		}

	}
	if len(toInt)%2 == 0 {
		fmt.Println("pari")
		sum += toInt[len(toInt)-1]

	} else {
		fmt.Println("dispari")
	}
	return sum
}
