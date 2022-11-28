package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := LeggiNumeri()
	fmt.Println(arr)

	var counted []int

	max := 0

	for _, i := range arr {
		if !find(counted, i) {
			//fmt.Println("Num: ", i)
			if Occorrenze(arr, i) == 1 {
				max += i

			}
			//fmt.Println(max)
			counted = append(counted, i)
		}
	}
	fmt.Println(max)

}

func LeggiNumeri() (numeri []int) {
	temp := os.Args
	for _, i := range temp {
		g, err := strconv.Atoi(i)
		if err == nil {
			numeri = append(numeri, g)
		}
	}
	return
}
func Occorrenze(numeri []int, n int) int {
	r := 0
	for _, i := range numeri {

		if i == n {
			r++
		}
	}
	return r
}

func find(numeri []int, n int) bool {
	for _, i := range numeri {
		if i == n {
			return true
		}
	}
	return false
}
