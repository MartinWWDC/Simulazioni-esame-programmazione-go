package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arrN := os.Args[2:]
	fmt.Println(arrN)
	EPSILON, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(EPSILON)
	for h := 0; h < len(arrN)-1; h++ {
		num1, _ := strconv.ParseFloat(arrN[h], 64)
		num2, _ := strconv.ParseFloat(arrN[h+1], 64)

		if num1-num2 > EPSILON {
			fmt.Print("<")
		} else if (num2 - num1) <= EPSILON {
			fmt.Print("=")
		} else if (num1 - num2) < (-EPSILON) {
			fmt.Print(">")
		}

	}

}
