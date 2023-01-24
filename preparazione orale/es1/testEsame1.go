package main

import "fmt"

/**
prendo una slice di interi e stampi la lunghezza massima sottosequenza di numeri interi
*/
func main() {
	arr := []int{0, 1, 2, 3, -4, 4, 3, 2, 8, 9, 19}
	fmt.Println(getMaxLenVtino(arr))
}

func getMaxLenVtino(s []int) int {
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
	nLen = append(nLen, len(s)-1-g)

	//fmt.Println(nLen)
	return getMax(nLen)
}
func getMax(arr []int) int {
	var max int
	for i := range arr {
		if i == 0 {
			max = arr[i]
		} else {
			if max < arr[i] {
				max = arr[i]
			}
		}
	}
	return max
}
func getMaxLenVProf(s []int) int {
	max := 0
	g := 0

	for i := range s {
		if !(s[i] >= 0) {
			if max < i-g {
				max = i - g
			}

			g = i

		}

	}

	//fmt.Println(nLen)
	return max
}
