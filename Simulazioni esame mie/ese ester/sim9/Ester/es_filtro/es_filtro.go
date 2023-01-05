package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var min, max int
	min, _ = strconv.Atoi(os.Args[1])
	max, _ = strconv.Atoi(os.Args[2])
	sum := 0
	for i := min + 1; i < max; i++ {
		if i%2 != 0 {
			//fmt.Println(i)
			sum += i

		}
	}
	fmt.Println(sum)
}
