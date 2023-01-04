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

	arr0 := strings.Split(eq, "x")
	var op int
	for h := range arr0[0] {
		if arr0[0][h] == '+' || arr0[0][h] == '-' {
			op = h
		}
	}

	c, _ := strconv.Atoi(arr0[0][:op])
	fmt.Println(c)
	b, _ := strconv.Atoi(arr0[0][op:])
	fmt.Println(b)
	a, _ := strconv.Atoi(arr0[1])
	fmt.Println(a)
	aPow, _ := strconv.Atoi(arr0[2][1 : len(arr0[2])-2])
	fmt.Println(aPow)
	x := (float64(b*(-1)) + (math.Sqrt(float64(b*b - 4*a*c)))) / float64(2*a)
	fmt.Println(x)
	x *= 100
	x = math.Round(x)
	x /= 100
	fmt.Println(x)

	y := (float64(b*(-1)) - (math.Sqrt(float64(b*b - 4*a*c)))) / float64(2*a)
	fmt.Println(y)
	y *= 100
	y = math.Round(y)
	y /= 100
	fmt.Println(y)

}
