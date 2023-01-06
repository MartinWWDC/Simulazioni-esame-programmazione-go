package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str1 := os.Args[1]
	str2 := os.Args[2]
	eq1, m1 := split(str1)
	eq2, m2 := split(str2)
	fmt.Println(eq1)
	fmt.Println(eq2)
	epsilon := 0.0001
	if m1-m2 <= epsilon {
		fmt.Println("le due rette sono parallele")
	} else {
		fmt.Println("le due rette non sono parallele")
	}

}

func split(str string) (arr []int, m float64) {
	arr1 := strings.Split(str, "x")
	arr2 := strings.Split(arr1[1], "y")
	arr3 := strings.Split(arr2[1], "=")
	g, _ := strconv.Atoi(arr1[0])
	arr = append(arr, g)
	g, _ = strconv.Atoi(arr2[0])
	arr = append(arr, g)
	g, _ = strconv.Atoi(arr3[0])
	arr = append(arr, g)
	m = -(float64(arr[0]) / float64(arr[1]))
	return
}
