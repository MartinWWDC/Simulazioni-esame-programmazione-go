package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

func main() {
	nc, _ := strconv.Atoi(os.Args[1])

	var num string
	var nPre, nPost string
	c := true
	fmt.Scanln(&num)
	for _, i := range num {
		if unicode.IsNumber(i) {
			if c {
				nPre += string(i)
			} else {
				nPost += string(i)

			}
		} else {
			if i == '.' {
				c = false
			}
		}
	}
	n, _ := strconv.Atoi(nPost)
	h := float64(n) / math.Pow(10, float64(len(nPost)))
	nP, _ := strconv.ParseFloat(nPre, 64)
	h *= math.Pow(10, float64(nc))

	h = math.Round(h)
	h = h / math.Pow(10, float64(nc))

	fmt.Println(nP + h)

}
