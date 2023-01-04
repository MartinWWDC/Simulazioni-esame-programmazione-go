package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var arr []float64
	var epsilon float64
	for i := 1; i <= len(os.Args)-1; i++ {
		g, _ := strconv.ParseFloat(os.Args[i], 64)
		if i == len(os.Args)-1 {

			epsilon = g
		} else {
			arr = append(arr, g)
		}
	}

	fmt.Println(arr)
	for i, g := range arr {
		if i != 0 {
			if arr[i-1] > g+epsilon {
				fmt.Print("<")
			} else if arr[i-1] < g-epsilon {
				fmt.Print(">")
			} else {
				fmt.Print("=")
			}
		}
	}

	//fmt.Println(epsilon)
}
