package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n1 := 1
	for _, n := range os.Args {
		if n2, err := strconv.Atoi(n); err == nil {
			n1 *= n2
		}
	}

	fmt.Println("Il risultato della moltiplicazione tra i numeri interi è ", n1)
}
