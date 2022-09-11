package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	g, ep := leggiInput()
	fmt.Println(convert(g, ep))
}

func convert(arr []float64, EPSILON float64) []string {
	arrf := []string{}
	for i := range arr {
		if i != len(arr)-1 {
			num1 := arr[i]
			num2 := arr[i+1]
			if num1-num2 > EPSILON {
				arrf = append(arrf, "<")
			} else if (num2 - num1) <= EPSILON {
				arrf = append(arrf, "=")
			} else if (num1 - num2) < (-EPSILON) {
				arrf = append(arrf, ">")
			}
		}

	}
	return arrf

}
func leggiInput() ([]float64, float64) {
	var arr []float64
	epsilon, _ := strconv.ParseFloat(os.Args[1], 64)
	for i := 2; i < len(os.Args); i++ {
		g, _ := strconv.ParseFloat(os.Args[i], 64)
		arr = append(arr, g)
	}
	return arr, epsilon
}
