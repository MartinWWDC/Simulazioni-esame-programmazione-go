package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	GeneraPrimi(os.Args[1])
}

func ÈPrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func GeneraPrimi(n string) (primi []int) {
	num, _ := strconv.Atoi(n)
	if ÈPrimo(num) {
		fmt.Println(num)
	}
	for lunghezza := 1; lunghezza <= 3; lunghezza++ {
		for pos := range n[:len(n)-lunghezza+1] {
			sottostringa := n[:pos] + n[pos+lunghezza:]
			n, _ := strconv.Atoi(sottostringa)
			if ÈPrimo(n) {
				fmt.Println(n)
				primi = append(primi, n)
			}
		}
	}
	return
}
