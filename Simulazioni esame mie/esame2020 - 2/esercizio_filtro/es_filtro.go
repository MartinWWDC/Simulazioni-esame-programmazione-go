package main

import "fmt"

func main() {
	var n, m float64
	fmt.Scan(&n, &m)
	bmi := n / (m * m)
	fmt.Println(bmi)
	if bmi >= 18.5 && bmi < 25 {
		fmt.Println("normale")
	} else if bmi < 18.5 {
		fmt.Println("sotto")
	} else {
		fmt.Println("sopra")
	}

}
