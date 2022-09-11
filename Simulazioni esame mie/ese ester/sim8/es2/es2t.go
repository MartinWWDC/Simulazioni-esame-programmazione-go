package main

import "fmt"

func main() {
	var str string
	str = "(())()"
	if isBalanced(str) {
		fmt.Println("bilanciata")
	} else {
		fmt.Println("non bilanciata")

	}
	fmt.Println(genStr(str))

}

func isBalanced(str string) bool {
	check := true
	openC := 0
	//closeC := 0
	for _, g := range str {
		if g != '(' && g != ')' {
			fmt.Println("h")
			check = false
		} else {

			if g == '(' {
				openC++

			} else {
				openC--

			}
		}
		if openC < 0 {
			//fmt.Println("c")
			check = false
		}
	}

	if openC != 0 {
		check = false
	}

	return check
}

func genStr(str string) map[string]int {
	mapp := make(map[string]int)
	for i := 0; i <= len(str); i++ {
		for g := i + 1; g <= len(str); g++ {
			if len(str[i:g]) > 1 {
				if isBalanced(str[i:g]) {
					//fmt.Println(str[i:g])
					mapp[str[i:g]]++
				}
			}
		}
	}
	return mapp
}
