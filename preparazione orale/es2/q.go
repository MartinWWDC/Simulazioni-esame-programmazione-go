package main

import "fmt"

/**
 * "rappresenta un struttura che rappresenta un float e un intero. attaccaci un metodo che incrementa il valore dell'intero "
 */
type tipo struct {
	n int
	f float64
}

func (re *tipo) IncreaseInt() {
	re.n++
}

func main() {
	t := tipo{0, 0.0}
	fmt.Println(t)
	t.IncreaseInt()
	fmt.Println(t)
}
