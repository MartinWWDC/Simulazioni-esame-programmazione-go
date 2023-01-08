package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N string
	var k int
	N = os.Args[1]
	k, _ = strconv.Atoi(os.Args[2])
	arr := strings.Split(N, "")
	maxs := getMax(arr, k)
	for j := 0; j < len(arr); j++ {
		u, _ := strconv.Atoi(arr[j])
		if !isIn(u, maxs) {
			fmt.Print(u)
		}
	}
}

func getMax(n []string, maxN int) []int {
	var maxAF []int
	for i := 0; i < maxN; i++ {
		max := 0
		for _, j := range n {
			h, _ := strconv.Atoi(j)
			if !isIn(h, maxAF) && max < h {
				max = h
			}
		}
		maxAF = append(maxAF, max)
	}
	//fmt.Println(maxAF)
	return maxAF
}

func isIn(n int, arr []int) bool {
	for _, i := range arr {
		if n == i {
			return true
		}
	}
	return false
}
