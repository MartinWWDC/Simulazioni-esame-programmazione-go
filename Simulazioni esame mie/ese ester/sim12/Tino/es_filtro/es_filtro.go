package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := os.Args[1:]
	str := ""
	for i := 0; i < len(arr); i += 2 {
		n, _ := strconv.Atoi(arr[i+1])
		for j := 0; j < n; j++ {
			str += arr[i]
		}

	}
	fmt.Println(str)
}
