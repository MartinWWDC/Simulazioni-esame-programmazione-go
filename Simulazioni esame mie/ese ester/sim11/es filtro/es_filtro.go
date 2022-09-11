package main

import (
	"fmt"
	"os"
)

func main() {
	str := os.Args[1]
	c := os.Args[2]
	count := 0
	for _, g := range str {
		if string(g) == (c) {
			count++
		}
	}
	fmt.Println(count)
}
