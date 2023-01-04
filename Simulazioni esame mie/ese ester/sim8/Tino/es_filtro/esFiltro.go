package main

import (
	"fmt"
)

func main() {
	var wrd string

	//fmt.Scanln(&wrd)
	wrd = "a1bG93foEìr92L"

	for h, g := range wrd {

		if h != len(wrd)-1 && (g >= 'a' && g <= 'z' || g >= 'A' && g <= 'Z') && (wrd[h+1] >= '1' && wrd[h+1] <= '9') {
			fmt.Println(string(wrd[h]))
		}

	}
}
