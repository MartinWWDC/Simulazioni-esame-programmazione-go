package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := []int{}
	n := os.Args[1]
	for j := range n {
		for h := j; h <= len(n); h++ {
			num, _ := strconv.Atoi(n[j:h])
			//fmt.Println(num)
			if num <= 100 && èPrimo(num) && isIn(arr, num) {
				fmt.Print(num, " ")
				arr = append(arr, num)
			}
		}
	}
}

func èPrimo(n int) bool {
	for h := 1; h < n; h++ {
		if n%h == 0 && h != 1 && h != n {
			return false
		}
	}
	if n == 1 || n == 0 {
		return false
	}
	return true
}
func isIn(n []int, c int) bool {
	for _, h := range n {
		if h == c {
			return false
		}
	}
	return true
}
