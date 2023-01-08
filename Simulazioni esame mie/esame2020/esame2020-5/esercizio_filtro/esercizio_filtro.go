package main

import (
	"fmt"
	"math"
)

func main() {
	var peso, alt float64
	fmt.Scan(&peso, &alt)
	bmi := peso / math.Pow(alt, 2)
	fmt.Println(bmi)
	if bmi >= 18.5 && bmi < 25 {
		fmt.Println("normale")
	} else {
		if bmi < 18.5 {
			fmt.Println("sotto")

		} else {
			fmt.Println("sopra")

		}
	}

}
