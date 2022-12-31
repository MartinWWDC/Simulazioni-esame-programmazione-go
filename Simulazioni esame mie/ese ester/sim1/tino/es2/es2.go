package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	divMin, _ := strconv.Atoi(os.Args[2])
	for i := 1; i <= n; i++ {
		for g := 1; g <= n; g++ {
			ndivI := trovaiDivisori(i)
			ndivG := trovaiDivisori(g)
			if areValid(ndivI, ndivG) && len(ndivG) >= divMin && isPerfect(i, ndivI) {
				fmt.Println(i, " ", g)
			}

		}
	}

}
func trovaiDivisori(n int) (arr []int) {
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			arr = append(arr, i)
		}
	}
	return
}
func areValid(arr1 []int, arr2 []int) bool {
	j := 0
	for _, i := range arr1 {
		for _, g := range arr2 {
			if i == g {
				j++
				break

			}
		}
	}
	return j <= 3
}
func isPerfect(n int, arr []int) bool {
	h := 0
	for _, i := range arr {
		h += i
	}
	return n == h
}
