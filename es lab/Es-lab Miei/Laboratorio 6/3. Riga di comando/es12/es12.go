package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, max, min := LeggiNumero()
	fmt.Println(n, max, min)

	(getDiv(n, min, max))
	//test(6, 12, 3, 4)

}

func test(i, j, min, max int) {
	divI := DivisoriPropri(i)
	divJ := DivisoriPropri(j)
	fmt.Println(divI, divJ)
	fmt.Println(isPerfect(divI, i), isPerfect(divJ, j))
	fmt.Println(len(divI))
	fmt.Println(len(divJ))
	if len(divJ) >= min {
		fmt.Println("check")
	}
	if len(divI) >= min && len(divJ) >= min {
		if len(divI) <= max && len(divJ) <= max && (isPerfect(divJ, j) || isPerfect(divI, i)) {
			fmt.Println(i, j)
		}
	}

}

func DivisoriPropri(n int) (arr []int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			arr = append(arr, i)
		}
	}
	return
}
func LeggiNumero() (n, min, max int) {
	str := os.Args[1]

	g, _ := strconv.Atoi(str)
	n = g
	str = os.Args[2]

	g, _ = strconv.Atoi(str)
	max = g
	str = os.Args[3]

	g, _ = strconv.Atoi(str)
	min = g
	return
}

func getDiv(n, min, max int) {
	var divI, divJ []int
	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			divI = DivisoriPropri(i)
			divJ = DivisoriPropri(j)
			divisoriComuni := DivisoriComuni(divI, divJ)

			if divisoriComuni <= max && divisoriComuni >= min && (isPerfect(divJ, j) || isPerfect(divI, i)) {
				fmt.Println(i, j)
			}
		}

	}
}
func isPerfect(num []int, n int) bool {
	i := 0
	var g int
	for _, g = range num {
		i += g
	}

	if i == n {
		return true
	} else {
		return false
	}
}

func DivisoriComuni(divisori1, divisori2 []int) (divisoriComuni int) {
	for _, divisore1 := range divisori1 {
		for _, divisore2 := range divisori2 {
			if divisore1 == divisore2 {
				divisoriComuni++
			}
		}
	}
	return
}
