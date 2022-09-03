package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := LeggiNumeri()
	m := Calcola(n)
	fmt.Println("La somma è", m)
}

func LeggiNumeri() []int {
	var narr []int
	for _, g := range os.Args[1:] {
		if n, err := strconv.Atoi(g); err == nil {
			narr = append(narr, n)
		}
	}
	return narr
}

func Calcola(sl []int) int {
	var y int
	end := 0

	for i, h := range sl {
		for g := i + 1; g < len(sl); g++ {
			//fmt.Println("(", h, ",", sl[g], ")")
			y = h * sl[g]
			if y%2 == 0 {
				end += y
			}
		}
	}
	return end
}
