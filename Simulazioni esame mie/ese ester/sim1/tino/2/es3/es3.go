package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Distretti struct {
	Nome     string
	Prefisso string
}
type Abbonato struct {
	cognome           string
	nome              string
	città             string
	cap               string
	via               string
	numero_civico     string
	numero_telefonico string
}

func main() {
	//fmt.Println(LeggiAbbonati())
	genMap()
}
func LeggiDistretti() (dis []Distretti) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		dis = append(dis, Distretti{arr[0], arr[1]})

	}
	return dis
}
func LeggiAbbonati() (abb []Abbonato) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		abb = append(abb, Abbonato{arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[5]})
	}
	return
}

func getDistetto(a Abbonato, di []Distretti) Distretti {
	for _, i := range di {
		if i.Nome == a.città {
			return i
		}
	}
	return Distretti{}
}
func genMap() {
	abb := LeggiAbbonati()
	dis := LeggiDistretti()
	fmt.Println(abb)
	fmt.Println(dis)
	mapp := make(map[Distretti]int)
	for _, a := range abb {
		d := getDistetto(a, dis)
		if d.Nome != "" {
			mapp[d]++
		}
	}
	fmt.Println(mapp)
}
