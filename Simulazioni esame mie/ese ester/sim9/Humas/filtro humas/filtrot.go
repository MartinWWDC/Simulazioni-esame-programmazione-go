package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	count := 0
	for i := n1 + 1; i < n2; i++ {
		if i%2 != 0 {
			count += i
		}
	}
	fmt.Println(count)
}
