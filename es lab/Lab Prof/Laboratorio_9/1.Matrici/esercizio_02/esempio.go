package main

import (
	"fmt"
)

func main() {

	var a [][]int

	fmt.Printf("Tipo di a: %T: %v\n", a, a)

	a = [][]int{{1, 2}, {10, 20}, {100, 200}}

	fmt.Printf("Tipo di a: %T: %v\n", a, a)

	for i := 0; i < len(a); i++ {
	//for i,_ := range a { // o si può usare il valore invece di a[i]
		fmt.Printf("a[%d] - tipo: %T: %v\n", i, a[i], a[i])
		for j := 0; j < len(a[i]); j++ {
		//for j,_ := range a[i] { // o si può usare il valore invece di a[i][j]
			fmt.Printf("a[%d][%d] - tipo: %T: %v\n", i, j, a[i][j], a[i][j])
		}
	}
}