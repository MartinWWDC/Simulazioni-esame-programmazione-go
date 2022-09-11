package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1234
	k := 1
	d := 4
	g := (GeneraNumeri(n, k))
	h := (FiltraNumeri(g, d))
	fmt.Println(h)
}

func GeneraNumeri(N, k int) (arr []int) {
	str := strconv.Itoa(N)
	for g := range str[:len(str)-k+1] {
		n, _ := strconv.Atoi(str[:g] + str[g+k:])
		arr = append(arr, n)
	}
	return arr
}

func FiltraNumeri(sl []int, d int) (arr []int) {
	for _, g := range sl {
		if g%d != 0 {
			arr = append(arr, g)
		}
	}
	return
}
