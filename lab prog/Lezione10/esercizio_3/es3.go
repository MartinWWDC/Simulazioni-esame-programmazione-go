package main

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]
	mapp := make(map[string]int)
	for i := range arr {
		for h := i; h <= len(arr); h++ {
			pto := arr[i:h]
			if len(pto) >= 3 && pto[0] == pto[len(pto)-1] {
				concat := ""
				for _, j := range pto {
					concat += j
				}
				mapp[concat]++
			}
		}
	}
	fmt.Println(mapp)
}
