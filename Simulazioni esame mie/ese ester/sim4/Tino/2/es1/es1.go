package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)

	for i := n1; i <= n2; i++ {
		c := true
		if i%3 == 0 {
			fmt.Print("Din")
			c = false

		}
		if i%5 == 0 {
			fmt.Print("Don")
			c = false

		}
		if i%7 == 0 {
			fmt.Print("Dan")
			c = false

		}
		if c {
			fmt.Print(i)
		}
		fmt.Print(" ")
	}

}
