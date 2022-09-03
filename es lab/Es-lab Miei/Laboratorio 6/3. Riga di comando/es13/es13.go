package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := LeggiNumeri()
	fmt.Println(arr)
	if len(arr) != 0 {
		fmt.Println(sum(arr))
	} else {
		fmt.Println("nopw")
	}

}

func LeggiNumeri() (arr []int) {
	if len(os.Args[1:])%2 == 0 {
		for _, g := range os.Args[1:] {
			if g, err := strconv.Atoi(g); err == nil {
				arr = append(arr, g)
			}
		}
	}
	return
}

func sum(arr []int) int {
	g := 0
	for i := 0; i < len(arr); i = i + 2 {
		g += arr[i] * arr[i+1]
	}
	return g
}
