package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var s int64
	var m1, q1, m2, q2, m3, q3, x, y float64
	fmt.Print("Inserisci s: ")
	fmt.Scan(&s)
	fmt.Print("Inserisci m1 e q1: ")
	fmt.Scan(&m1, &q1)
	fmt.Print("Inserisci m2 e q2: ")
	fmt.Scan(&m2, &q2)
	fmt.Print("Inserisci m3 e q3: ")
	fmt.Scan(&m3, &q3)
	s = 3
	m1 = 0
	q1 = 2
	m2 = -1
	q2 = 15
	m3 = 1
	q3 = 0
	rand.Seed(s)
	for i := 0; i < 10; i++ {
		x = rand.Float64() * 10
		y = rand.Float64() * 10
		fmt.Print("Punto(", x, ",", y, ")")
		check(m1, q1, m2, q2, m3, q3, x, y)
	}

}

func check(m1, q1, m2, q2, m3, q3, v1, v2 float64) {
	//if ÈXUgualeAY(v1, v2*m1+q1) || ÈXUgualeAY(v1, v2*m2+q2) || ÈXUgualeAY(v1, v2*m3+q3) {
	//	fmt.Println("è iterno")

	if ÈXMinoreDiY(v2, v1*m1+q1) || ÈXMaggioreDiY(v2, v1*m2+q2) || ÈXMaggioreDiY(v2, v1*m3+q3) {
		fmt.Println("è esterno ")
	} else {
		fmt.Println("è iterno")
	}
}

const EPSILON = 1.0e-9

func ÈXMaggioreDiY(x, y float64) bool {
	return (x - y) > EPSILON
}

func ÈXUgualeAY(x, y float64) bool {
	return math.Abs(x-y) <= EPSILON
}

func ÈXMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}
