package main

import "fmt"

func main() {
	fmt.Println(conta(1111, 1))
}

func conta(i, count int) int {
	if i%10 == count {
		count = conta(i/10, count+1)
		return count
	} else {
		return conta(i/10, count)

	}
}
