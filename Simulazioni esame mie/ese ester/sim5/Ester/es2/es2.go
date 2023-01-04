package main

import "fmt"

func main() {
	g := LegggiStringa()
	fmt.Println(g)
	fmt.Println(Sottostringhe(g))
}

func LegggiStringa() string {
	var str string
	check := true
	fmt.Scan(&str)
	for _, g := range str {
		if g >= '1' && g <= '9' {

		} else {
			fmt.Println("err:", g)
			check = false
		}
	}
	if check {
		return str
	} else {
		return ""
	}
}

func Sottostringhe(s string) map[string]int {
	mapp := make(map[string]int)
	var check bool
	for i := 0; i < len(s); i++ {
		//g := len(s)
		for g := len(s) - 2; g >= i; g++ {

			str := s[i:g]
			fmt.Println(str)
			if len(str) > 1 {
				for j := 0; j < len(str); j++ {
					check = true
					if j != 0 && str[j-1] < str[j] {
						check = false
					}
				}
				if check {
					mapp[str]++
				}
			}

		}
	}

	return mapp
}
