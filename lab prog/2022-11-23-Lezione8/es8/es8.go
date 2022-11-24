package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	nums := LeggiNumeri()
	suf, nons := FiltraVoti(nums)
	fmt.Println("sufi:", suf)
	fmt.Println("non:", nons)

}
func LeggiNumeri() (numeri []int) {
	temp := os.Args[1:]
	for _, n := range temp {
		h, err := strconv.Atoi(n)
		if err == nil {
			numeri = append(numeri, h)

		}
	}
	return
}

func FiltraVoti(voti []int) (sufficienti, insufficienti []int) {

	for _, i := range voti {
		if i < 60 {
			insufficienti = append(insufficienti, i)
		} else {
			sufficienti = append(sufficienti, i)
		}
	}
	return
}
