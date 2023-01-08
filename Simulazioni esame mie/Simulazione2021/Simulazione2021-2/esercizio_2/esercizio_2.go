package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	//str = "abcdef"
	fmt.Println(str)
	//arr := strings.Split(str, "")
	check := true
	for _, g := range str {
		if g >= 'a' && g <= 'z' {

		} else {
			check = false
		}
	}

	if check {

		fmt.Println(SottoStringe(str))
	}
}

func SottoStringe(str string) map[string]int {
	check := true
	mappa := make(map[string]int)
	for i := 0; i <= len(str); i++ {
		for j := i; j <= len(str); j++ {
			strin := str[i:j]
			if len(strin) > 1 {
				for k := 0; k < len(strin)-1; k++ {
					if k < len(strin) && strin[k] < strin[k+1] {

					} else {
						check = false
					}

				}
				if check {
					mappa[strin]++
				}
				check = true
			}
		}
	}
	return mappa
}
