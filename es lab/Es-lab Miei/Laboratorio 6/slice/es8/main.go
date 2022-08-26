/*
Scrivere un programma che,

dopo aver letto da standard input un numero intero n,

chiede all'utente di inserire n numeri interi (sempre da standard input).

Il programma deve stampare gli n numeri interi in ordine inverso rispetto a quello di inserimento.

Suggerimento: Per creare dinamicamente una slice si utilizzi la funzione make().

*/
package main

import (
	"fmt"
)

func main() {
	var n int
	var tmp float64
	var arr []float64

	fmt.Print("N:")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&tmp)
		arr = append(arr, tmp)

	}

	inverso := make([]float64, len(arr))
	for i := 0; i < n; i++ {
		inverso[i] = arr[len(arr)-i-1]
	}
	fmt.Println(inverso)

}
