package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	txt := LeggiTesto()
	//txt := "testo di prova \n su cui provare le statistiche"
	nline, nwords, lenghtmedia := (StatisticheParole(txt))
	fmt.Println(nline)
	fmt.Println(nwords)
	fmt.Println(float64(lenghtmedia) / float64(nwords))
}

func LeggiTesto() (text string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}
	return
}

func StatisticheParole(s string) (numerolinee int, numparole int, lunghezzaTotale int) {
	lines := strings.Split(s, "\n")
	for _, li := range lines {
		if len(li) > 0 {
			numerolinee++
		}
		words := strings.Split(li, " ")
		numparole += len(words)
		for _, g := range words {
			lunghezzaTotale += len(g)
		}
	}
	numparole--
	return
}
