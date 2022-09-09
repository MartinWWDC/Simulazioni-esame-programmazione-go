package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fmt.Print(check(28, 42, 4))
	test()
}

func test() {
	n, _ := strconv.Atoi(os.Args[1])
	divmin, _ := strconv.Atoi(os.Args[2])
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {

			if check(i, j, divmin) {
				fmt.Println(i, j)
			}

		}
	}
}
func check(i, j, divmin int) bool {
	var arri, arrj []int
	for g := 1; g < i; g++ {
		if i%g == 0 {
			arri = append(arri, g)
		}
	}
	for g := 1; g < j; g++ {
		if j%g == 0 {
			arrj = append(arrj, g)
		}
	}
	//fmt.Println(arri, arrj)
	count := 0
	//fmt.Println(isPerfect(28, arri))
	//mappa = make(map[string]int)

	if len(arrj) >= divmin && isPerfect(i, arri) == true {
		for g := 0; g < len(arri); g++ {

			for h := 0; h < len(arrj); h++ {

				if arri[g] == arrj[h] {

					count++
					//fmt.Println(count)
				}

			}

		}
		if count > 3 {
			return false
		}

	} else {
		return false
	}
	return true

}

func isPerfect(i int, arri []int) bool {
	h := 0
	for g := 0; g < len(arri); g++ {
		//fmt.Print(h)
		h += arri[g]
	}
	if i == h {
		return true
	} else {
		return false
	}
}
