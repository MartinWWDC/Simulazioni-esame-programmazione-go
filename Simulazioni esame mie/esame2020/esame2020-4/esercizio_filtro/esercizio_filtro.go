package main

import "fmt"

func main() {
	var peso, altezza float64
	fmt.Scan(&peso)
	fmt.Scan(&altezza)
	//fmt.Println(peso, " ", altezza)
	bmi := (peso / (altezza * altezza))
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
