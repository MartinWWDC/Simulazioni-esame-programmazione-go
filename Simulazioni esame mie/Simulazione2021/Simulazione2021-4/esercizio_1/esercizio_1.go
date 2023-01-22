package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	eq := os.Args[1]
	arr := strings.Split(eq, "x")
	sogli, _ := strconv.ParseFloat(os.Args[2], 64)
	ep, _ := strconv.ParseFloat(os.Args[3], 64)
	fmt.Println(arr)
	a, _ := strconv.ParseFloat(arr[0], 64)
	b, _ := strconv.ParseFloat(arr[1][3:], 64)
	c, _ := strconv.ParseFloat(arr[2][:len(arr[2])-2], 64)
	fmt.Println(a, b, c)
	sol1 := (-1*b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	sol1 *= 1000
	sol1 = math.Round(sol1)
	sol1 /= 1000
	sol2 := (-1*b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	sol2 *= 1000
	sol2 = math.Round(sol2)
	sol2 /= 1000
	if sol1 == sol2 {
		fmt.Println("una sola sol: ", sol1)
	} else {
		fmt.Println("due sol:", sol1, sol2)
	}
	if (sol1)-ep > sogli {
		fmt.Println("sol1 magg:", sol1)
	}

	if (sol2)-ep > sogli {
		fmt.Println("sol1 magg:", sol2)
	}

}
