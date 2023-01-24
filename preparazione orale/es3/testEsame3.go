package main

import "fmt"

/**
 * Scrivere una funzione che prende una slice di float e restituisce un intero e deve calcolarte quanti positivi sono seguiti da un negativo
 */

func main() {
	arr := []float64{1.0, 2, 3, -4}
	positiviPostNeg(arr)
}

func positiviPostNeg(n []float64) int {
	g := 0
	for i := 0; i < len(n)-1; i++ {
		if n[i] > 0 && n[i+1] < 0 {
			fmt.Println(n[i])
			g++
		}
	}
	return g
}
