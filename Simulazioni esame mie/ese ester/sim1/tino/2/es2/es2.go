package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	N, _ := strconv.Atoi(os.Args[1])
	divMin, _ := strconv.Atoi(os.Args[2])

	for a := 1; a <= N; a++ {
		for b := 1; b <= N; b++ {
			if areValid(a, b, divMin) {
				fmt.Println("-", a, " ", b)
			}
		}
	}
	fmt.Println(areValid(28, 12, 4))
}
func DivisoriPropri(n int) []int {
	var arr []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

func areValid(a, b int, divMin int) bool {
	g := 0
	divA := DivisoriPropri(a)
	divB := DivisoriPropri(b)
	//fmt.Println(len(divB))
	if isPerfect(a) && len(divB) >= divMin {

		for i := 0; i < len(divA); i++ {
			if isIn(divB, divA[i]) {
				g++
			}
		}
		return g <= 3

	} else {
		//fmt.Println("2")
	}

	return false
}
func isIn(arr []int, n int) bool {
	for _, i := range arr {
		if i == n {
			return true
		}
	}
	return false
}
func isPerfect(n int) bool {
	arr := DivisoriPropri(n)
	//fmt.Println(arr)
	//fmt.Println(n)
	n2 := 0
	for _, i := range arr {
		n2 += i
		//fmt.Println(n2)
	}
	//fmt.Println(n2 == n)
	return n2 == n
}
