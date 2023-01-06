package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	mapp := make(map[string]int)
	arr := []string{}
	n := (os.Args[1])
	for h := range n {
		for i := h; i <= len(n); i++ {

			str := n[h:i]
			if len(str) >= 2 && isCresc(str) {
				arr = append(arr, str)
			}

		}
	}
	fmt.Println(arr)
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	fmt.Println(arr)
	for h := len(arr) - 1; h >= 0; h-- {

		mapp[arr[h]]++

	}

	fmt.Println(mapp)

}
func isCresc(n string) bool {
	for i := range n {
		if i != len(n)-1 && !(n[i] < n[i+1]) {
			return false
		}
	}
	return true
}
