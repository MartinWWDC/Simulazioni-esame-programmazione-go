package main

import . "fmt"

func isBalanced(s string) bool {
	var c int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			c++
		} else {
			c--
		}
		//Println(c)
		if c < 0 {
			return false
		}
	}
	if c != 0 {
		return false
	} else {
		return true
	}

}

func sott(s string) map[string]int {
	mappa := make(map[string]int)
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			sott := s[i:j]
			//Println(sott)
			if len(sott) > 1 && isBalanced(sott) {
				mappa[sott]++
			}
		}
	}
	return mappa
}

func main() {
	/*
		var s string
		Scan(&s)
	*/
	s := "())("
	for _, r := range s {
		if r != '(' && r != ')' {
			return
		}
	}
	if isBalanced(s) {
		Println("Bilanciata")
	} else {
		Println("Non Bilanciata")
	}
	mappa := sott(s)
	Println(mappa)
}
