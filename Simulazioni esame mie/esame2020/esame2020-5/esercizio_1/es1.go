package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := os.Args[1]
	fmt.Println(getSubs(s))
}

func getSubs(s string) map[string]int {
	arr := strings.Split(s, "")
	//fmt.Println(arr)
	mapp := make(map[string]int)
	for i := range arr {
		for h := i + 1; h < len(arr); h++ {
			str := concat(arr[i : h+1])

			if isPalindroma(str) {
				mapp[str]++
			}
		}
	}
	return mapp
}

func isPalindroma(s string) bool {
	sIn := ""
	arr := strings.Split(s, "")
	for i := len(arr) - 1; i >= 0; i-- {
		sIn += string(arr[i])
	}
	return sIn == s
}

func concat(s []string) (wrd string) {
	for _, j := range s {
		wrd += j
	}
	return
}
