package main

import (
	"bufio"
	"fmt"
	"math"
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
	//GetPercorso(ut, 106, 1)
	//GetPercorso(ut, 107, 0.65)
	//GetPercorso(ut, 68, 0.5)
	GetPercorso(ut, 67, 1)

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
	return "Percorso da " + StringPosizione(percorso[0]) + " a " + StringPosizione(percorso[len(percorso)-1]) + " cambi: " + strconv.Itoa(len(percorso))
}

func GetPercorso(pos []Posizione, distanza float64, soglia float64) {
	var dis float64
	i := 0
	for j := 1; j < len(pos); j++ {
		if pos[j].longitude < 0 && pos[j-1].longitude > 0 {
			dis = 180 - math.Abs(pos[j].longitude) + 180 - math.Abs(pos[j-1].longitude)
		} else if pos[j].longitude > 0 && pos[j-1].longitude < 0 {
			dis = math.Abs(pos[j-1].longitude) + pos[j].longitude
		} else if pos[j].longitude < 0 && pos[j-1].longitude < 0 {
			dis = math.Abs(math.Abs(pos[j].longitude) - math.Abs(pos[j-1].longitude))
		} else {
			dis = math.Abs(pos[j].longitude - pos[j-1].longitude)
		}
		if (math.Abs(pos[j].latitude-pos[j-1].latitude) > soglia || dis > distanza) && (j-i) != 1 {
			fmt.Println(StringPercorso(pos[i:j]))
			i = j
		} else if j == len(pos)-1 {
			fmt.Println(StringPercorso(pos[i : j+1]))
		}

	}
}
