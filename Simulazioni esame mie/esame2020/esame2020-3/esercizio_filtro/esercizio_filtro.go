package main

import "fmt"

func main() {
	var p, h float64
	fmt.Scan(&p)
	fmt.Scan(&h)
	bmi := p / (h * h)
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
