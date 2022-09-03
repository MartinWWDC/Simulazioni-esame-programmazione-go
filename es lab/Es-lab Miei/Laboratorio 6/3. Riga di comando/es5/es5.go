package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := LeggiNumeri()
	fmt.Println(Minimo(arr))
	fmt.Println(Massimo(arr))
	fmt.Println(Media(arr))
}

func LeggiNumeri() (numeri []int) {
	arr := os.Args[1:]
	var arr2 []int
	for _, el := range arr {
		if n, err := strconv.Atoi(el); err == nil {
			arr2 = append(arr2, n)
		}
	}
	return arr2
}

func Minimo(sl []int) int {
	var min int
	for i, el := range sl {
		if i == 0 {
			min = el
		} else {
			if min > el {
				min = el
			}
		}
	}
	return min
}

func Massimo(sl []int) int {
	var max int
	for i, el := range sl {
		if i == 0 {
			max = el
		} else {
			if max < el {
				max = el
			}
		}
	}
	return max
}

func Media(sl []int) float64 {
	media := 0
	for _, n := range sl {
		media += n
	}
	return (float64(media) / float64(len(sl)))
}
