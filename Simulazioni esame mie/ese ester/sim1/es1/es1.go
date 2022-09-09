package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var n int
	n, _ = strconv.Atoi(os.Args[1])
	var str string

	fmt.Scan(&str)
	//str = "&4&$4%mamma5lsmia.6cosa1suc0ced8"
	strfloat := ""
	for _, g := range str {
		if (g >= '0' && g <= '9') || g == '.' {
			fmt.Println(string(g))

			strfloat += string(g)
		}
	}

	float, _ := strconv.ParseFloat(strfloat, 64)
	fmt.Println(float)
	fmt.Println(RoundFloat(float, n))

}

func RoundFloat(number float64, decimalPlace int) float64 {
	temp := math.Pow(10, float64(decimalPlace))
	return math.Round(number*temp) / temp
}
