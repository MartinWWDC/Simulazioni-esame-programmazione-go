package main

import "fmt"

func main() {
	var n, g int
	var arr []int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&g)
		arr = append(arr, g)
	}
	for i := n - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
}
