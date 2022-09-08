package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	soglia, _ := strconv.ParseFloat(os.Args[2], 64)
	epsisoln, _ := strconv.ParseFloat(os.Args[3], 64)

	estrai_a := strings.Split(os.Args[1], "x^2")
	a, _ := strconv.ParseFloat(estrai_a[0], 64)

	estrai_b := strings.Split(estrai_a[1], "x")
	b, _ := strconv.ParseFloat(estrai_b[0], 64)

	estrai_c := strings.Split(estrai_b[1], "=")
	c, _ := strconv.ParseFloat(estrai_c[0], 64)
	fmt.Println(a, b, c)
	discriminante := b*b - (4 * a * c)

	sol1 := (-b + math.Sqrt(discriminante)) / (2 * a)
	sol2 := (-b - math.Sqrt(discriminante)) / (2 * a)
	if sol1 == sol2 {
		fmt.Println("un’unica soluzione reale:", sol1)
	} else {
		fmt.Println(sol2)

		fmt.Println(sol1)
	}

	if sol1-soglia > epsisoln {
		fmt.Println("valore ", sol1, " è maggiore della soglia ")

	}
	if sol2-soglia > epsisoln {
		fmt.Println("valore ", sol2, " è maggiore della soglia ")

	}

}
func Split(r rune) bool {
	return r == '+' || r == '-'
}
