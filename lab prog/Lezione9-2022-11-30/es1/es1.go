package main

import (
	"fmt"
	"math/rand"
)

type CarteDaGioco struct {
	tipo    string
	simbolo string
}
type Mazzo struct {
	carte  [40]CarteDaGioco
	nCarte int
}

func main() {
	mazzo := CreaMazzo()
	fmt.Println(mazzo)
	mazzo = shuffle(mazzo)
	fmt.Println(mazzo)
}

func CreaCarta(seme, simbolo string) CarteDaGioco {
	return CarteDaGioco{seme, simbolo}
}

func CreaMazzo() Mazzo {
	var mazzo Mazzo
	semi := []string{"quadri", "picche", "cuori", "fiori"}
	numeri := []string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette", "fante",
		"donna", "re"}
	h := 0
	for _, i := range semi {
		for _, j := range numeri {
			mazzo.carte[h] = CreaCarta(i, j)
			h++
		}
	}
	mazzo.nCarte = 40
	return mazzo

}
func shuffle(mazzo Mazzo) Mazzo {
	max := rand.Intn(100)
	for i := 0; i < max; i++ {
		h := rand.Intn(40)
		r := rand.Intn(40)
		for h == r {
			h = rand.Intn(40)
			r = rand.Intn(40)
			fmt.Println("uguali")
		}
		mazzo = switchCards(mazzo, h, r)
	}
	return mazzo
}
func switchCards(mazzo Mazzo, i int, j int) Mazzo {
	h := mazzo.carte[i]
	mazzo.carte[i] = mazzo.carte[j]
	mazzo.carte[j] = h
	return mazzo
}
func Preleva(m Mazzo) (CarteDaGioco, Mazzo, error) {
	carta := m.carte[m.nCarte-1]
	m.carte[m.nCarte-1] = CreaCarta("", "")
	m.nCarte -= 1
	return carta, m, nil
}
func Riponi(m Mazzo, c CarteDaGioco) (Mazzo, error) {
	m.carte[m.nCarte] = c
	m.nCarte++
	return m, nil
}
