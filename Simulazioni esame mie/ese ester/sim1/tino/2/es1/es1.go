package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var str string
	fmt.Scanln(&str)
	var n int
	n, _ = strconv.Atoi(os.Args[1])
	strM := ""
	for _, i := range str {

		if (i >= '0' && i <= '9') || i == '.' {
			strM += (string(i))
		}
	}
	fmt.Println(strM)
	k, _ := strconv.ParseFloat(strM, 64)
	k *= math.Pow(10, float64(n))
	k = math.Round(k)
	k /= math.Pow(10, float64(n))
	fmt.Println(k)

}
