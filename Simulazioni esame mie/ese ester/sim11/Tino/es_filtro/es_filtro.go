package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	wrd := os.Args[1]
	char := os.Args[2]
	arr := strings.Split(wrd, "")
	c := 0
	for _, h := range arr {
		if h == char {
			c++
		}
	}
	fmt.Println(c)
}
