package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fun := os.Args[1]
	soglia, _ := strconv.ParseFloat(os.Args[2], 64)
	epsilon, _ := strconv.ParseFloat(os.Args[3], 64)
	arr := strings.Split(fun, "x")
	arr[1] = arr[1][2:]
	arr[2] = strings.Split(arr[2], "=")[0]
	fmt.Println(arr)
	a, _ := strconv.ParseFloat(arr[0], 64)
	b, _ := strconv.ParseFloat(arr[1], 64)
	c, _ := strconv.ParseFloat(arr[2], 64)
	xP := (-1*b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	xM := (-1*b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	fmt.Println(xP, ";", xM)
	if xP-soglia > epsilon {
		fmt.Println("La soluzione: ", xP, " è maggiore della soglia  ")
	}
	if xM-soglia > epsilon {
		fmt.Println("La soluzione: ", xM, " è maggiore della soglia  ")
	}
}
