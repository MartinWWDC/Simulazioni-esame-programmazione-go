package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]

	arr := strings.Split(str, "")
	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)/2 {
			fmt.Println(str)
		} else {
			for j := 0; j < len(arr)/2; j++ {
				fmt.Print(" ")
			}
			fmt.Println(arr[i])
		}
	}
}
