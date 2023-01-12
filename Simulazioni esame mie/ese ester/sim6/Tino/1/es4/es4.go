package main

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]
	getSubs(arr)

}

func getSubs(arr []string) {
	mapp := make(map[string]int)
	for i := range arr {
		for j := i + 2; j < len(arr); j++ {
			if arr[i] == arr[j] {
				var str string
				for _, h := range arr[i : j+1] {
					str += h
				}
				mapp[str]++

			}
		}
	}
	fmt.Println(mapp)
}
