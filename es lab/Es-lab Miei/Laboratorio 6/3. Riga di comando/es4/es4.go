package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var narr []int
	arr := os.Args[1:]
	for _, n := range arr {
		if num, err := strconv.Atoi(n); err == nil {
			narr = append(narr, num)
		}
	}
	fmt.Println("Minimo: ", Minimo(narr))
	fmt.Println("Massimo: ", Massimo(narr))
	fmt.Println("Media: ", Media(narr))

}

func Minimo(sl []int) int {
	var min int
	for i, s := range sl {
		if i == 0 {
			min = s
		} else {
			if s < min {
				min = s
			}
		}
	}
	return min
}
func Massimo(sl []int) int {
	var max int
	for i, s := range sl {
		if i == 0 {
			max = s
		} else {
			if s > max {
				max = s
			}
		}
	}
	return max
}
func Media(sl []int) float64 {
	var media int
	for _, s := range sl {
		media += s
	}
	return float64(media) / float64(len(sl))

}
