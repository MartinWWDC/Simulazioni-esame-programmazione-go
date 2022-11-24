package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	nums := LeggiNumeri()
	fmt.Println(nums)
	re := Calcola(nums)
	fmt.Println(re)
}

func Calcola(sl []int) (mul int) {
	mul = 0
	for g := range sl {
		for j := g; j < len(sl); j++ {
			if j != g && (sl[g]*sl[j])%2 == 0 {
				fmt.Println("g:", sl[g], " j:", sl[j])

				mul += sl[g] * sl[j]

			}

		}
	}
	return
}
func LeggiNumeri() (numeri []int) {
	temp := os.Args[1:]
	for _, n := range temp {
		h, err := strconv.Atoi(n)
		if err == nil {
			numeri = append(numeri, h)

		}
	}
	return
}
