package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N, k int
	fmt.Scan(&N, &k)
	str := strconv.Itoa(N)
	for h := range str[:len(str)-k+1] {
		g := str[:h] + str[h+k:]
		fmt.Println(g)

	}
}
