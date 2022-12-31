package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, L, H int
	//fmt.Scan(&n)
	//fmt.Scan(&L)
	//fmt.Scan(&H)
	n, _ = strconv.Atoi(os.Args[1])
	L, _ = strconv.Atoi(os.Args[2])
	H, _ = strconv.Atoi(os.Args[3])
	fmt.Println(n, L, H)
	count := 0
	c := true
	for i := 0; i < n+n+1; i++ {
		if i%2 == 0 {
			if c {
				for h := 0; h < L; h++ {
					fmt.Print(count)
					count = counter(count)
				}
				fmt.Println()
			} else {
				arr := []int{}
				for h := 0; h < L; h++ {
					//fmt.Print(count)
					arr = append(arr, count)
					count = counter(count)
				}
				for h := len(arr) - 1; h >= 0; h-- {
					fmt.Print(arr[h])
				}
				fmt.Println()
			}

		} else {
			for j := 0; j < H-2; j++ {
				if j != 0 {
					fmt.Println()
				}
				for h := 0; h < L; h++ {

					if c {
						if h == L-1 {
							fmt.Print(count)
							count = counter(count)

						} else {
							fmt.Print(" ")
						}
					} else {
						if h == 0 {
							fmt.Print(count)
							count = counter(count)

						} else {
							fmt.Print(" ")
						}
					}

					//fmt.Print(count)
					//count = counter(count)
				}

			}
			if c {
				c = false
			} else {
				c = true
			}
			if H-2 != 0 {
				fmt.Println()
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
