package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	genSub("1234567890")
}
func genSub(n string) {
	mapp := make(map[int]int)
	for i := 0; i < len(n); i++ {
		for g := i + 1; g <= len(n); g++ {
			str := n[i:g]
			n, _ := strconv.Atoi(str)
			if n <= 100 && IsPrime(n) {
				mapp[n]++
			}
		}
	}
	for h, _ := range mapp {
		fmt.Println(h)
	}

}
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
