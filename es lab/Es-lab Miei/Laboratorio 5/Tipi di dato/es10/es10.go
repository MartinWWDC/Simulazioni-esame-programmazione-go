package main

import (
	"fmt"
	"math"
	"math/rand"
)

const EPSILON = 1.0e-9

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è maggiore
di `y` di almeno una costante EPSILON */
/* ÈXMaggioreDiY(5.0,4.999) -> true */
/* ÈXMaggioreDiY(5.0,4.9999999999) -> false */
func ÈXMaggioreDiY(x, y float64) bool {
	return (x - y) > EPSILON
}

/* La funzione `func ÈXUgualeAY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è uguale
a `y` a meno di una costante EPSILON */
/* ÈXUgualeAY(5.0,4.999) -> false */
/* ÈXUgualeAY(5.0,4.9999999999) -> true */
func ÈXUgualeAY(x, y float64) bool {
	return math.Abs(x-y) <= EPSILON
}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è minore
di `y` di almeno una costante EPSILON */
/* ÈXMinoreDiY(4.999,5.0) -> true */
/* ÈXMinoreDiY(4.9999999999,5.0) -> false */
func ÈXMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}

func main() {
	var s, m, q int64
	var x, y float64
	fmt.Print("Inserisci s: ")
	fmt.Scan(&s)
	fmt.Print("Inserisci m e q: ")
	fmt.Scan(&m, &q)
	fmt.Println(s, m, q)
	rand.Seed(s)
	for i := 0; i < 10; i++ {
		x = rand.Float64() * 5
		y = rand.Float64() * 5
		fmt.Print("Punto (", x, ",", y, ") -")
		if ÈXMaggioreDiY(y, x) {
			fmt.Println("Il punto sta sopra la retta")
		} else if ÈXMinoreDiY(y, x) {
			fmt.Println("Il punto sta sotto la retta")
		} else {
			fmt.Println("Il punto appartiene alla retta")

		}
	}

	fmt.Println(x, y)
}
