package main

import (
	"fmt"
)

func main() {
	var n int
	var nums []string
	var tmp string
	fmt.Scanln(&n)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		nums = append(nums, tmp)
	}

	fmt.Println(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Print(nums[i], " ")
	}
}
