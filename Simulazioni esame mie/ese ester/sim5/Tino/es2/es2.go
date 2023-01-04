package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	LeggiInput()

}
func isDecres(arr []string) bool {

	for h := len(arr) - 1; h > 0; h-- {

		a, _ := strconv.Atoi(arr[h-1])

		b, _ := strconv.Atoi(arr[h])

		if a <= b {

			return false

		}

	}

	return true

}

func LeggiInput() {

	var n string
	//fmt.Scanln(&n)
	n = "654321"
	arr := strings.Split(n, "")
	arrN := []string{}
	for h := 0; h <= len(arr)-1; h++ {
		for i := h + 1; i <= len(arr); i++ {
			if len(arr[h:i]) > 1 {
				fmt.Println(arr[h:i])

				if !isDecres(arr[h:i]) {
					continue
				} else {
					//fmt.Println(arr[h:i])
					str := ""
					for _, j := range arr[h:i] {
						str += j
					}
					arrN = append(arrN, str)
				}
			}
		}
	}
	fmt.Println(arrN)
	sort.Strings(arrN)
	mapp := make(map[string]int)
	for _, h := range arrN {
		mapp[h]++
	}
	fmt.Println((mapp))
}
