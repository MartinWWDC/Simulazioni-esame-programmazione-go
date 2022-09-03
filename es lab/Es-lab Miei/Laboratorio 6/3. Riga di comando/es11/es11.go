package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	arr := gen(os.Args[1])
	fmt.Println(arr)
}

func gen(numero string) []int {
	var arr []int
	n, _ := strconv.Atoi(numero)

	if ÈPrimo(n) {
		fmt.Println(n)
	}

	for lunghezza := 1; lunghezza <= 3; lunghezza++ {
		for posizione := range numero[:len(numero)-lunghezza+1] {
			sottostringa := numero[:posizione] + numero[posizione+lunghezza:]
			n, _ := strconv.Atoi(sottostringa)
			fmt.Println(n)

			if ÈPrimo(n) {
				//fmt.Println(n)
				arr = append(arr, n)
			}
		}
	}
	fmt.Println()

	return arr
}

func ÈPrimo(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)

}
