package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	s := os.Args[1]
	mapp := (genSub(s))
	var arr []string
	fmt.Println(mapp)
	for i := range mapp {
		arr = append(arr, i)
	}
	sort.SliceStable((arr), func(i int, j int) bool { return len(arr[i]) > len(arr[j]) })
	for _, i := range arr {
		fmt.Println(i, mapp[i])
	}
}

func genSub(str string) map[string]int {
	mappa := make(map[string]int)
	for i := 0; i <= len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			s := str[i:j]
			if check(s) {
				mappa[s]++
			}

		}
	}
	return mappa
}

func check(str string) bool {
	return str[0] == str[len(str)-1] && len(str) >= 3
}
