package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	mul := 1
	temp := os.Args[1:]
	for _, n := range temp {
		if j, err := strconv.Atoi(n); err == nil {
			mul *= j
		}
	}
	fmt.Println(mul)
}
