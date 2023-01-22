package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	c := 1
	o := true
	for i := 0; i <= (2*n)-1; i++ {
		for j := 0; j < (2*n)-1; j++ {
			if j <= n-1 && i < n {
				if i == 0 || j == n-1 {
					fmt.Print("0")
				} else {
					if j == c && i == c {
						fmt.Print("0")
						c++
					} else {
						fmt.Print("*")
					}
				}
			} else if j >= n-1 && i > n-1 {
				if o {
					//c++
					o = false
				}
				if i == 2*n-1 || j == n-1 {
					fmt.Print("0")

				} else {
					if j == c && i == c || j == n+1 {
						fmt.Print("0")
						c++
					} else {
						fmt.Print("*")
					}
				}
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println("")
	}
}
