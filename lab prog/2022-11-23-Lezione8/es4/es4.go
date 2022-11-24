package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum, j := 0, 0
	var min, max int
	temp := os.Args[1:]
	fmt.Println(temp)
	for i, n := range temp {
		h, _ := strconv.Atoi(n)

		if h > max || i == 0 {
			max = h
		}
		if h < min || i == 0 {
			min = h
		}
		sum += h
		j++
	}
	fmt.Println(j)
	fmt.Println("Minimo: ", min)
	fmt.Println("Massimo: ", max)
	fmt.Println("medio: ", float64(sum)/float64(j))
}
