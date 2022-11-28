package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	h, _ := strconv.Atoi(os.Args[1])
	f := Fattoriali(h)
	fmt.Println(f)

}

func Fattoriali(n int) (arr []int) {
	for g := 1; g <= n; g++ {
		fmt.Println("g:", g)
		temp := 1
		for h := 1; h <= g; h++ {
			fmt.Println(h)

			temp *= h
		}
		arr = append(arr, temp)
	}
	return
}
