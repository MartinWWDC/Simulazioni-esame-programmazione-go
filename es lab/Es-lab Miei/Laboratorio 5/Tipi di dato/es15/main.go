/*
Scrivere un programma che legga da standard input un valore intero e sei valori reali:

il primo valore è il seme (seed) s da utilizzare per inizializzare il generatore di numeri casuali;
il secondo ed il terzo valore sono il coefficiente angolare m1 e il termine noto q1 di una retta r1: y = m1*x + q1 sul piano cartesiano su cui giace la base AB di un triangolo ABC;
il quarto ed il quinto valore sono il coefficiente angolare m2 e il termine noto q2 di una retta r2: y = m2*x + q2 sul piano cartesiano su cui giace il lato BC di un triangolo ABC;
il sesto ed il settimo valore sono il coefficiente angolare m3 e il termine noto q3 di una retta r3: y = m3*x + q3 sul piano cartesiano su cui giace il lato AC di un triangolo ABC.
Una volta terminata la fase di lettura, il programma deve generare per 10 volte una coppia di valori reali px e py che rappresentano rispettivamente l'ascissa e l'ordinata di un punto sul piano cartesiano e, per ciascun punto:

determinare se, a meno di una costante EPSILON = 1.0e-9, il punto sta all'interno o all'esterno del triangolo ABC;
stampare a video il relativo messaggio (come mostrato nell'Esempio di esecuzione).
I valori px e py devono essere compresi nell'intervallo [0, 10.0).
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var seed int
	counter := 10
	var (
		m1, q1 float64
		m2, q2 float64
		m3, q3 float64
	)
	seed = 3
	m1 = 0
	q1 = 2
	m2 = -1
	q2 = 15
	m3 = 1
	q3 = 0
	//fmt.Println("seed:")
	//fmt.Scanln(&seed)
	//fmt.Println("il coefficiente angolare m1 e il termine noto q1")
	//fmt.Scanln(&m1, &q1)
	//fmt.Println("il coefficiente angolare m2 e il termine noto q2")
	//fmt.Scanln(&m2, &q2)
	//fmt.Println("il coefficiente angolare m3 e il termine noto q3")

	//fmt.Scanln(&m3, &q3)
	bo := false
	var poinX []float64
	var poinY []float64
	rand.Seed(int64(seed))
	for i := 0; i < counter; i++ {
		poinX = append(poinX, rand.Float64()*10)
		poinY = append(poinY, rand.Float64()*10)
	}
	for i := 0; i < counter; i++ {
		fmt.Println("punto ", i, "(", poinX[i], ")(", poinX[i], ")")
		eq1 := m1*poinX[i] + q1
		eq2 := m2*poinX[i] + q2
		eq3 := m3*poinX[i] + q3
		switch {
		case eq1 == poinY[i] && eq2 == poinY[i]:
			bo = true
		case eq2 == poinY[i] && eq3 == poinY[i]:
			bo = true
		case eq1 == poinY[i] && eq3 == poinY[i]:
			bo = true

		}
		if bo {
			fmt.Println("si")
		}

	}

}
