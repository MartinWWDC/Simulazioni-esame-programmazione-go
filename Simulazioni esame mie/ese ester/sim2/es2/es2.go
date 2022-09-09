package main

import "fmt"

func main() {
	var n, L, H int
	//fmt.Scan(&n, &L, &H)
	n = 2
	L = 3
	H = 4
	count := 0
	//fmt.Println(n, L, H)

	for i := 0; i <= n+2; i++ {
		for g := 0; g < n-1; g++ {
			if i%2 == 0 {
				for h := 0; h < L-1; h++ {
					fmt.Print(count)
					count = counter(count)
				}
			} else {
				for y := 0; y < H-1; y++ {
					fmt.Println(count)
					count = counter(count)
				}
			}
		}
	}

}
func counter(i int) int {
	if i < 9 {
		i++

	} else {
		i = 0
	}
	return i
}
