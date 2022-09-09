package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var n, k, d int
	n := 1234
	k := 1
	d := 4
	g := GeneraNumeri(n, k)
	fmt.Println(g)
	fmt.Println(FiltraNumeri(g, d))
}

func GeneraNumeri(N, k int) (arr []int) {
	str := strconv.Itoa(N)
	for i := range str[:len(str)-k+1] {
		sub := str[:i] + str[i+k:]
		s, _ := strconv.Atoi(sub)
		arr = append(arr, s)

	}
	return
}

func FiltraNumeri(sl []int, d int) (arr []int) {
	for _, g := range sl {
		fmt.Println(g % d)
		if g%d != 0 {
			arr = append(arr, g)
		}
	}
	return
}
