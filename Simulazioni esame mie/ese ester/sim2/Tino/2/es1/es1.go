/**
 * dio cane il ruomuovere cifreeee
 * svegliaaaaaaaaaaaaaa
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := (getSubs())
	for _, i := range arr {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
func isPrime(n int) bool {
	for i := 1; i < n; i++ {
		if i != 1 && i != n && n%i == 0 {
			return false
		}
	}
	return true
}
func getSubs() (arr []int) {
	n := os.Args[1]
	for dis := 1; dis <= 3; dis++ {
		for i := range n[:len(n)-dis+1] {

			j, _ := strconv.Atoi(n[:i] + n[i+dis:])
			arr = append(arr, j)
		}
	}
	return
}
