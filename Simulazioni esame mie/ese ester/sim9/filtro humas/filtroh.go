package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	var s int
	for i := a + 1; i < b; i++ {
		if i%2 != 0 {
			s += i
		}
	}
	Println(s)
}
