package main

import (
	"fmt"
	"os"
)

func main() {
	arr := (LeggiInput())
	mapp := GetOcc(arr)
	fmt.Println(mapp)
}
func LeggiInput() (arr []string) {
	word := os.Args[1]
	for i := 0; i <= len(word); i++ {
		for g := i + 2; g < len(word); g++ {
			sub := word[i : g+1]
			if sub[0] == sub[len(sub)-1] {
				arr = append(arr, sub)
			}
		}
	}
	return
}

func GetOcc(arr []string) map[string]int {
	mapp := make(map[string]int)
	for _, h := range arr {
		mapp[h]++
	}
	return mapp
}
