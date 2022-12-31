package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Comando struct {
	direzione string
	passi     int
}

func main() {

	Min := Comando{"", 0}
	g := LeggiIndicazioni()
	fmt.Println(g)
	mapp := (AnalizzaComandi(g))
	counter := 0
	//invMapI := []int{}
	//invMapDir := []string{}
	//for _, i := range g {
	//	switch {
	//	case i.direzione == "NORD":
	//		invMapDir = append(invMapDir, "SUD")
	//	case i.direzione == "SUD":
	//		invMapDir = append(invMapDir, "NORD")
	//	case i.direzione == "OVEST":
	//		invMapDir = append(invMapDir, "OVEST")
	//	case i.direzione == "NORD":
	//		invMapDir = append(invMapDir, "EST")
	//	}
	//	invMapI = append(invMapI, i.passi)
	//
	//}
	for i, g := range mapp {

		//fmt.Println(i, Min.direzione, g, Min.passi)
		if counter == 0 {
			Min = Comando{i, g}
		} else {
			if Min.passi > g {
				Min = Comando{i, g}

			}
			if Min.direzione == i {
				arr := []string{Min.direzione, i}
				sort.Strings(arr)
				if Min.direzione != arr[0] {
					Min = Comando{i, g}
				}
			}
		}
		counter++

	}
	fmt.Println("Min: ", Min.direzione)
	fmt.Println("Comandi Inversi")
	for i := 0; i <= len(g)-1; i++ {
		h := g[len(g)-i-1]
		//fmt.Print(h.direzione, " ")
		switch {
		case h.direzione == "NORD":
			fmt.Print("SUD")
		case h.direzione == "SUD":
			fmt.Print("NORD")
		case h.direzione == "EST":
			fmt.Print("OVEST")
		case h.direzione == "OVEST":
			fmt.Print("EST")

		}
		fmt.Print(" ", g[len(g)-i-1].passi, " ")
	}
}

func LeggiIndicazioni() (lista []Comando) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		dir := arr[0]
		pass, _ := strconv.Atoi(arr[1])
		lista = append(lista, Comando{dir, pass})
	}
	return
}

func AnalizzaComandi(comandi []Comando) (mappa map[string]int) {
	mappa = make(map[string]int)
	for _, g := range comandi {
		mappa[g.direzione] += g.passi
	}
	return
}
