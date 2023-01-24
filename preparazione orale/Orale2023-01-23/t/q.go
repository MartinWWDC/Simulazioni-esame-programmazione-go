package main

import "fmt"

type tipo struct {
	f float64
	i int
}

func (t *tipo) RaddI() {
	t.i *= 2
}

func main() {
	v := tipo{0, 2}
	fmt.Println(v)

	v.RaddI()
	fmt.Println(v)
}
