package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	l := LeggiInput()
	//fmt.Println(l)
	fmt.Println(Sottostringhe(l))
	//test()
}
func test() {
	v1 := []int{1, 2, 3, 4}
	v2 := []int{1, 2, 3, 4}
	g := compare(v1, v2)
	fmt.Println(g)
}
func LeggiInput() string {
	check := true
	for _, g := range os.Args[1] {
		if g >= '0' && g <= '9' {

		} else {
			check = false
		}
	}
	if check {
		return os.Args[1]
	} else {
		return ""
	}
}
func Sottostringhe(s string) (subn []string) {
	var arri, temp []int
	for i := 0; i <= len(s); i++ {
		for j := i + 2; j <= len(s); j++ {
			arri = nil
			str := s[i:j]
			//fmt.Println(str)
			for _, g := range str {
				l, _ := strconv.Atoi(string(g))

				arri = append(arri, l)
			}
			//fmt.Println(arri)
			temp = clone(arri)
			sort.Ints(arri)
			//fmt.Println("sort", arri, temp, compare(arri, temp))

			if compare(arri, temp) {

				subn = append(subn, str)
			}
		}
	}
	return
}
func compare(n, s []int) bool {
	check := true
	if len(n) == len(s) {
		for i := 0; i < len(n); i++ {
			if n[i] != s[i] {
				check = false
			}
		}
	} else {
		check = false
	}
	return check
}
func clone(s []int) (arr []int) {
	for _, n := range s {
		arr = append(arr, n)
	}
	return
}
