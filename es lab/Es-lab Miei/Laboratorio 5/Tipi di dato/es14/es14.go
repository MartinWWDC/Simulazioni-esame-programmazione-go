package main

import (
	"fmt"
	"math"
)

func main() {
	var xA, yA float64
	var xB, yB float64
	var xC, yC float64
	var AB, BC, AC float64
	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto A:")
	fmt.Scan(&xA, &yA)
	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto B:")
	fmt.Scan(&xB, &yB)
	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto C:")
	fmt.Scan(&xC, &yC)
	AB = Distanza(xA, yA, xB, yB)
	BC = Distanza(xB, yB, xC, yC)
	AC = Distanza(xA, yA, xC, yC)

	if AB == BC && BC == AC {
		fmt.Print("Il triangolo ABC è equilatero.\n Lunghezza del lato: ")
	} else if AB == BC || BC == AC || AB == AC {
		fmt.Print("Il triangolo ABC è isocele.")

	} else {
		fmt.Print("Il triangolo ABC è Scaleno.")

	}
}

func Distanza(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}
