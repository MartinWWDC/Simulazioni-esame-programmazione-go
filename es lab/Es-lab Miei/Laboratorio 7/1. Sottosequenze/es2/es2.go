package main

import (
	"fmt"
	"os"
)

func main() {
	str, err := LetturaInput()
	fmt.Println(str, err)
	fmt.Println(isBalanced(str))
	fmt.Println(subSet(str))
}

func LetturaInput() (string, bool) {
	check := true
	for _, g := range os.Args[1] {
		//fmt.Println(string(g), rune('('))
		if g != rune('(') && g != rune(')') {
			check = false
		}
	}
	return os.Args[1], check

}

func isBalanced(str string) bool {
	counter := 0

	for _, g := range str {
		if g == rune('(') {
			counter++
		}
		if g == rune(')') {
			counter--
		}
	}
	return counter == 0
}
func subSet(str string) (arr []string) {
	for i, g := range str {
		if g == rune('(') {
			arr = append(arr, genSub(str, i))
		}
	}
	arr = append(arr, str)
	return
}

func genSub(str string, start int) string {
	for i := start; i < len(str); i++ {
		if str[i] == ')' && isBalanced(str[start:i+1]) {
			return str[start : i+1]

		}
	}
	return ""
}
