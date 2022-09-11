package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var str string
	//fmt.Scan(&str)
	str = "123456"

	mapp := genSub(str)
	fmt.Println(mapp)
	fmt.Println(len(mapp))
	var slice []int
	for i := range mapp {
		t, _ := strconv.Atoi(i)
		slice = append(slice, t)
	}
	sort.Ints(slice)
	fmt.Println(slice)
	for i := len(slice) - 1; i >= 0; i-- {
		t := strconv.Itoa(slice[i])
		fmt.Println(slice[i], mapp[t])
	}

}

func genSub(str string) map[string]int {
	mapp := make(map[string]int)
	var check bool
	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			s := str[i:j]
			if len(s) > 1 {
				check = true
				for h := 0; h < len(s); h++ {
					if h != len(s)-1 {
						if s[h] < s[h+1] {

						} else {
							check = false
						}
					}
				}
				if check {
					mapp[s]++
				}
			}

		}
	}
	return mapp
}
