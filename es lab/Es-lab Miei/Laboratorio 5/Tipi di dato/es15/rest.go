package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(55)
	fmt.Println(rand.Float64())
}
