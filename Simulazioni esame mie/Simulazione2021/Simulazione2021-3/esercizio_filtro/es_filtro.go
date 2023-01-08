package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := (os.Args[1])
	arr := strings.Split(n, "")
	for i := 0; i < len(arr)-1; i++ {
		n, _ := strconv.Atoi(arr[i])
		nS, _ := strconv.Atoi(arr[i+1])
		if nS > n {
			fmt.Print(n)
		}

	}

}
