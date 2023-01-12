package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Posizione struct {
	nome      string
	latitude  float64
	longitude float64
}

func main() {
	ut := LeggiUtenze()
	for _, u := range ut {
		fmt.Println(StringPosizione(u))
	}
	fmt.Println(StringPercorso(ut))
}
func LeggiUtenze() []Posizione {
	var pos []Posizione
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		x, _ := strconv.ParseFloat(arr[1], 64)
		y, _ := strconv.ParseFloat(arr[2], 64)
		p, err := NuovaPosizione(arr[0], x, y)
		if !err {
			pos = append(pos, p)
		}
	}
	return pos
}
func NuovaPosizione(etichetta string, latitudine float64, longitudine float64) (Posizione, bool) {
	if !(latitudine >= -90 && latitudine <= 90) {
		return Posizione{}, true
	}
	if !(longitudine >= -180 && longitudine <= 180) {
		return Posizione{}, true
	}
	return Posizione{etichetta, latitudine, longitudine}, false
}
func StringPosizione(posizione Posizione) string {
	return posizione.nome + " (" + fmt.Sprint(posizione.latitude) + "," + fmt.Sprint(posizione.longitude) + ")"
}

func StringPercorso(percorso []Posizione) string {
	return "Percorso da " + StringPosizione(percorso[0]) + " a " + StringPosizione(percorso[len(percorso)-1]) + " cambi: " + strconv.Itoa(len(percorso)-2)
}

func GetPercorso(pos []Posizione, distanza float64, soglia float64) {

}
