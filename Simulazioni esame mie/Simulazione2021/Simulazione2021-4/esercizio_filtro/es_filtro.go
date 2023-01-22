package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := os.Args[1]
	for i := 1; i < len(n); i++ {
		a, _ := strconv.Atoi(string(n[i-1]))
		b, _ := strconv.Atoi(string(n[i]))
		if a < b {
			fmt.Print(a)
		}
	}
}
