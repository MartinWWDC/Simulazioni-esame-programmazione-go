package main

import "fmt"

/**
prendo una slice di interi e stampi la lunghezza massima sottosequenza di numeri interi
*/
func main() {
	arr := []int{0, 1, 2, 3, -4, 4, 3}
	fmt.Println(getMaxLen(arr))
}

func getMaxLen(s []int) int {
	nLen := []int{}
	g := 0

	for i := range s {
		if s[i] >= 0 {
		} else {
			if i != g {
				nLen = append(nLen, i-g)
			}
			g = i

		}

	}
	fmt.Println(nLen)
	return g
}
