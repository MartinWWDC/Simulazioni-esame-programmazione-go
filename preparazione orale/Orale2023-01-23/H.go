package main

type tipo struct {
	f float64
	i int
}

func (t *tipo) Adder() {
	t.f++
}
