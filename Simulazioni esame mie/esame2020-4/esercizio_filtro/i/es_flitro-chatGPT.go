package main

import "fmt"

func main() {
	// Dichiariamo le variabili per il peso e l'altezza
	var weight, height float64

	// Leggiamo il peso e l'altezza da standard input utilizzando fmt.Scan
	fmt.Scan(&weight, &height)

	// Calcoliamo il BMI
	bmi := weight / (height * height)

	// Stampiamo il BMI e se è nella norma o meno
	fmt.Printf("Il tuo BMI è: %.2f\n", bmi)
	if bmi >= 18.5 && bmi < 25 {
		fmt.Println("Il tuo peso è nella norma.")
	} else if bmi < 18.5 {
		fmt.Println("Il tuo peso è al di sotto della norma.")
	} else {
		fmt.Println("Il tuo peso è al di sopra della norma.")
	}
}
