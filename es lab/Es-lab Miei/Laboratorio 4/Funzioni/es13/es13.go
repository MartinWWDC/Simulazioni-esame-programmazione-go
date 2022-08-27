package main

import (
	"fmt"
	"math"
)

func main() {
	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	if ÈPrimo(soglia) {
		fmt.Print("Numeri primi inferiori a")
		fmt.Print(soglia)
		fmt.Println(": ")
		NumeriPrimi(soglia)

	}

}
func ÈPrimo(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	if n > 0 && n%n == 0 && n%1 == 0 {
		if n < 2 {
			fmt.Println("Number must be greater than 2.")
			return false
		}
		for i := 2; i <= sqrt; i++ {
			if n%i == 0 {
				//fmt.Println("Non Prime Number")
				return false
			}
		}
		//fmt.Println("Prime Number")
		return true
	} else {
		return false
	}

}

func NumeriPrimi(limite int) {
	for i := 2; i < limite; i++ {
		if ÈPrimo(i) == true {
			fmt.Print(i)
		}

	}
}
