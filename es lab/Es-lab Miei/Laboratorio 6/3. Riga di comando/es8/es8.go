package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := LeggiNumeri()
	suf, ins := FiltraVoti(n)
	fmt.Println(suf)
	fmt.Println(ins)
}

func LeggiNumeri() (numeri []int) {
	var narr []int
	for _, g := range os.Args[1:] {
		if n, err := strconv.Atoi(g); err == nil {
			narr = append(narr, n)
		}
	}
	return narr

}

func FiltraVoti(voti []int) (sufficienti, insufficienti []int) {
	for _, n := range voti {
		if n >= 60 {
			sufficienti = append(sufficienti, n)
		} else {
			insufficienti = append(insufficienti, n)
		}
	}
	return
}
