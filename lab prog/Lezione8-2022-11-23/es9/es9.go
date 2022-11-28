package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	soglia, _ := strconv.Atoi(os.Args[1])
	fmt.Println(Genera(soglia))

}

func Genera(soglia int) (nums []int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	numeroGenerato := rand.Intn(100)
	for numeroGenerato < soglia {
		fmt.Println(numeroGenerato)

		nums = append(nums, numeroGenerato)
		numeroGenerato = rand.Intn(100)

	}
	return

}
