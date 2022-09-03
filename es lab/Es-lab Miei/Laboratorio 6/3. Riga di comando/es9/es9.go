package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	g := Genera(n)
	fmt.Println(g)
	fmt.Println(g[:len(g)-1])
}

func Genera(soglia int) []int {
	rand.Seed(int64(time.Now().Nanosecond()))
	check := true
	var arr []int
	var numeroGenerato int
	for check {
		numeroGenerato = rand.Intn(100)
		if numeroGenerato >= soglia {
			arr = append(arr, numeroGenerato)
		} else {
			arr = append(arr, numeroGenerato)
			check = false

		}

	}
	return arr
}
