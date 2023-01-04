package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var eq = "-1-10x-4x^2=0"
	arr := strings.Split(eq, "=")
	arr2 := strings.Split(arr[0], "x")
	fmt.Println(arr2)
	arr3 := strings.Split(arr2[0], "+")
	fmt.Println(arr3)

	//fmt.Println(arr3[1])
	fmt.Println(arr2[0])
	t := 0
	for i := 1; i < len(arr2[0]); i++ {
		//fmt.Println(string(arr2[0][i]))
		if arr2[0][i] == '-' || arr2[0][i] == '+' {
			t = i
		}
	}
	//c := arr2[0][0:t]
	c, _ := strconv.Atoi(arr2[0][0:t])
	//b := arr2[0][t:]
	b, _ := strconv.Atoi(arr2[0][t:])
	a, _ := strconv.Atoi(arr2[len(arr2)-2])
	fmt.Println(a, b, c)
	res1 := (-float64(b) + math.Sqrt(float64(b*b-4*a*c))) / float64(2*a)
	res2 := (-float64(b) - math.Sqrt(float64(b*b-4*a*c))) / float64(2*a)
	res1 *= 100
	res2 *= 100
	res1 = math.Round(res1) / 100
	res2 = math.Round(res2) / 100

	fmt.Println(res1, res2)
}
