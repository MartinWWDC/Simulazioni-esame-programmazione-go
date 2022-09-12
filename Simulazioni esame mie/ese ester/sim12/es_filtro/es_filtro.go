package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var final string
	//var str string
	for i := 1; i < len(os.Args); i++ {
		if i != len(os.Args)-1 {
			char := os.Args[i]
			n, _ := strconv.Atoi(os.Args[i+1])
			for g := 0; g < n; g++ {
				final += char
			}

		}
	}
	fmt.Println(final)
}
