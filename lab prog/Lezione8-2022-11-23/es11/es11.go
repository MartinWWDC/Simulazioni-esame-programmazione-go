package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := LeggiNumeri()
	h := Calcola(arr)
	fmt.Println(h)

}
func LeggiNumeri() (numeri []int) {
	temp := os.Args
	for _, i := range temp {
		g, err := strconv.Atoi(i)
		if err == nil {
			numeri = append(numeri, g)
		}
	}
	return
}
func Calcola(sl []int) int {
	i, h := 0, 0
	for i < len(sl) {
		h += sl[i] * sl[i+1]
		i += 2
	}
	return h
}
