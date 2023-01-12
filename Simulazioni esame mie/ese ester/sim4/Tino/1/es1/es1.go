package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	for i := n1; i <= n2; i++ {
		arrD := getDivisori(i)
		c := true
		for _, g := range arrD {
			switch g {
			case 3:
				fmt.Print("Din")
				c = false
			case 5:
				fmt.Print("Don")
				c = false

			case 7:
				fmt.Print("Dan")
				c = false

			}
		}
		fmt.Print(" ")
		if c {
			fmt.Print(i)
			fmt.Print(" ")
		}

	}

}

func getDivisori(n int) (arrDiv []int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			arrDiv = append(arrDiv, i)
		}
	}
	return
}
