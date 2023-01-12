package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := os.Args[1:]
	epsilon, _ := strconv.ParseFloat(arr[len(arr)-1], 64)
	fmt.Println("epp:", epsilon)
	for h := 1; h < len(arr)-1; h++ {
		last, _ := strconv.ParseFloat(arr[h-1], 64)
		main, _ := strconv.ParseFloat(arr[h], 64)
		ep := main - last
		//fmt.Println(last, main)
		//fmt.Println(ep)
		switch {
		case ep >= -epsilon && ep <= epsilon:
			fmt.Print("=")
		case ep > epsilon:
			fmt.Print(">")
		case ep < epsilon:
			fmt.Print("<")

		}
	}
}
