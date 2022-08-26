/*
Scrivere un programma che legga da standard input sei valori reali:

il primo ed il secondo valore, xA e yA, rappresentano rispettivamente l'ascissa e l'ordinata di un punto A sul piano cartesiano;
il terzo ed il quarto valore, xB e yB, rappresentano rispettivamente l'ascissa e l'ordinata di un punto B sul piano cartesiano;
il quinto ed il sesto valore, xC e yC, rappresentano rispettivamente l'ascissa e l'ordinata di un punto C sul piano cartesiano.
Una volta terminata la fase di lettura, il programma deve stampare a video (come mostrato nell'Esempio di esecuzione), se il triangolo ABC definito dai segmenti/lati AB, BC e AC è equilatero, iscoscele o scaleno.

Un triangolo è equilatero se ha tutti e tre i lati di lunghezza uguale.

Un triangolo è isoscele se ha solo due lati di lunghezza uguale.

Un triangolo è scaleno se ha tutti e tre i lati di lunghezza diversa.

La lunghezza di ciascun lato di un triangolo è pari alla distanza euclidea tra gli estremi del lato.

Per esempio, la lunghezza del lato AB del triangolo ABC è pari alla distanza euclidea tra i punti A e B: ((xA-xB)2 + (yA-yB)2)1/2.

Le lunghezze dei lati del triangolo vanno confrontate a meno di una costante EPSILON = 1.0e-9.

Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione Distanza(x1, y1, x2, y2 float64) float64 che riceve in input:

due valori float64 nei parametri x1 e y1 che rappresentano rispettivamente l'ascissa e l'ordinata di un punto P1 sul piano cartesiano;

due valori float64 nei parametri x2 e y2 che rappresentano rispettivamente l'ascissa e l'ordinata di un punto P2 sul piano cartesiano;

e restituisce un valore float64 pari alla distanza euclidea tra i punti P1 e P2.

*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		xA, yA float64
		xB, yB float64
		xC, yC float64
	)

	//xA = 1
	//yA = 1
	//xB = 1
	//yB = 4
	//xC = 1.5
	//yC = 1.5
	fmt.Print("Cordinate x,y punto A")
	fmt.Scanln(&xA, &yA)
	fmt.Print("Cordinate x,y punto B")
	fmt.Scanln(&xB, &yB)
	fmt.Print("Cordinate x,y punto C")
	fmt.Scanln(&xC, &yC)
	AB := calcolaDIstanza(xA, yA, xB, yB)
	BC := calcolaDIstanza(xB, yB, xC, yC)
	CA := calcolaDIstanza(xC, yC, xA, yA)
	fmt.Println(checkType(AB, BC, CA))

}
func calcolaDIstanza(xA float64, yA float64, xB float64, yB float64) float64 {
	distance := math.Sqrt((xA-xB)*(xA-xB) + (yB-yA)*(yB-yA))
	return distance
}
func checkType(AB, BC, CA float64) string {
	var s string
	fmt.Println("controllo lati AB,BC:", checkEqual(AB, BC))
	fmt.Println("controllo lati BC,CA:", checkEqual(BC, CA))
	fmt.Println("controllo lati AB,CA:", checkEqual(AB, CA))
	switch {

	case checkEqual(AB, BC) == true && checkEqual(BC, CA) == true && checkEqual(AB, CA) == true:
		s = "equilatero\n Lato="

	case checkEqual(AB, BC) != true && checkEqual(BC, CA) != true && checkEqual(AB, CA) != true:
		s = "scaleno"

	case checkEqual(AB, BC) == checkEqual(BC, CA) || checkEqual(AB, BC) == checkEqual(AB, CA) || checkEqual(BC, CA) == checkEqual(BC, CA):
		s = "isoscele"

	}

	return s

}
func checkEqual(a, b float64) bool {
	if a == b {
		return true
	} else {
		return false
	}
}
