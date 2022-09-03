package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	narr := LeggiNumeri()
	for g := range narr {
		if Occorrenze(narr, g) == 1 {
			sum += g
		}
	}
	fmt.Println(sum)
}

func LeggiNumeri() (numeri []int) {
	for _, g := range os.Args[1:] {
		if n, err := strconv.Atoi(g); err == nil {
			numeri = append(numeri, n)
		}
	}
	return numeri
}
func Occorrenze(numeri []int, n int) int {
	i := 0
	for _, g := range numeri {
		if n == g {
			i++
		}
	}
	return i
}
