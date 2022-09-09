package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Posizione struct {
	et   string
	lat  float64
	long float64
}
type percorso []Posizione

func main() {
	g := LeggiPercorso()
	fmt.Println(g)
	fmt.Println(StringPercorso(g))
}
func LeggiPercorso() (perc percorso) {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		et := arr[0]
		lat, _ := strconv.ParseFloat(arr[1], 64)
		long, _ := strconv.ParseFloat(arr[2], 64)
		pos, _ := NuovaPosizione(et, lat, long)
		perc = append(perc, pos)
	}
	return
}

func NuovaPosizione(etichetta string, latitudine float64, longitudine float64) (Posizione, bool) {
	check := false
	if latitudine >= -90 && latitudine <= 90 && longitudine > -180 && longitudine <= 180 {
		check = true

	}
	return Posizione{etichetta, latitudine, longitudine}, check
}

func StringPosizione(posizione Posizione) string {
	return posizione.et + "(" + fmt.Sprint(posizione.lat) + "," + fmt.Sprint(posizione.long)
}
func StringPercorso(percorso []Posizione) (str string) {
	cambi := 0
	distanza, _ := strconv.ParseFloat(os.Args[1], 64)
	soglia, _ := strconv.ParseFloat(os.Args[2], 64)
	for i := 0; i < len(percorso)-1; i++ {
		for j := i; j < len(percorso); j++ {
			h := percorso[i:j]
			prev := 0
			if len(h) > 2 {
				for p := 1; p < len(h)-1; p++ {
					if h[prev].long-h[p].long < distanza && (h[prev].lat-h[p].lat) < soglia {
						cambi++
						prev = p
					} else {
						str += "da" + StringPosizione(h[prev]) + StringPosizione(h[p]) + string(cambi) + "\n"
					}
				}
			}
		}
	}
	return
}
