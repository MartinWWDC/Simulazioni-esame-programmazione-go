package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	soglia, _ := strconv.Atoi(os.Args[1])

}

func Genera(soglia int) (nums []int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for numeroGenerato := rand.Intn(100); numeroGenerato<soglia{
		nums = append(nums,numeroGenerato)
	}
	return

}
