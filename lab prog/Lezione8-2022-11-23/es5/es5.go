package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	h := LeggiNumeri()
	fmt.Println(h)
}
func LeggiNumeri() (numeri []int) {
	temp := os.Args[1:]
	for _, n := range temp {
		h, err := strconv.Atoi(n)
		if err == nil {
			numeri = append(numeri, h)

		}
	}
	return
}

// ridicolo, il mio è meglio
