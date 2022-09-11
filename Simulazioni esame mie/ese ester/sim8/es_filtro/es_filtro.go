package main

import "fmt"

func main() {
	var str string

	fmt.Scan(&str)
	for i, g := range str {
		if i != len(str)-1 && (g >= 'a' && g <= 'z' || g >= 'A' && g <= 'Z') && (str[i+1] >= '1' && str[i+1] <= '9') {
			fmt.Println(string(g))
		}
	}
}
