package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Scan(&str)
	//str = "sottotono"
	arr := strings.Split(str, "")
	mapp := make(map[string]int)
	c := true
	for i := range arr {
		for h := i; h <= len(arr); h++ {
			str := arr[i:h]
			if len(str) > 1 {
				//fmt.Println(str, len(str))

				for k := range str {
					//fmt.Println(str[k], "==", str[len(str)-k-1])
					if str[k] == str[len(str)-k-1] {

					} else {
						//fmt.Println("skip")
						c = false
					}

				}
				if c {
					g := ""
					for i := range str {
						g += str[i]
					}
					mapp[g]++
				} else {
					c = true
				}

			}
		}
	}
	for i := range mapp {
		fmt.Println(i, ": ", mapp[i])
	}

}
