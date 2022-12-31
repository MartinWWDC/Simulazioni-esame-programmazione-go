package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type posizione struct {
	etc  string
	lat  float64
	long float64
}

func main() {
	posizioni := []posizione{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		//fmt.Println(arr)
		lat, _ := strconv.ParseFloat(arr[1], 64)
		long, _ := strconv.ParseFloat(arr[2], 64)
		pos, err := nuovaPosizione(arr[0], lat, long)
		if !err {
			//fmt.Println("add")
			posizioni = append(posizioni, pos)

		}
	}
	/*fmt.Println(posizioni)
	for _, j := range posizioni {
		fmt.Println(StringPosizione(j))
	}*/
	Check(posizioni, 1, 106)
}
func nuovaPosizione(etichetta string, lat float64, long float64) (posizione, bool) {
	if lat >= -90.00 && lat <= 90.00 && long >= -180.00 && long <= 180.00 {
		return posizione{etichetta, lat, long}, false

	} else {
		return posizione{"", 0, 0}, true

	}
}

func StringPosizione(posizione posizione) string {
	return posizione.etc + "(" + fmt.Sprint(posizione.lat) + "," + fmt.Sprint(posizione.long) + ")"
}

func Check(pos []posizione, dist float64, soglia float64) {
	n := 0
	u := 0.0
	k := 0.0
	fmt.Print("Percorso da ", StringPosizione(pos[0]), " a ")
	main := pos[0]
	for j := 1; j < len(pos); j++ {

		fmt.Println(pos[j].long, "-", pos[j-1].long, pos[j].long-pos[j-1].long)
		//u += pos[j-1].long - pos[j].long
		//k += pos[j].lat - pos[j-1].lat
		if pos[j-1].long <= pos[j].long && pos[j-1].long-pos[j].long <= dist && pos[j-1].lat-pos[j].lat <= soglia {
			n++
		} else {
			fmt.Print(StringPosizione(pos[j]), " spostamenti: ", n)
			fmt.Println("")
			fmt.Println("n:", u, k)
			if j < len(pos)-1 {
				main = pos[j+1]
				n = 0
				u = 0
				k = 0
				fmt.Println()
				fmt.Print("Percorso da ", StringPosizione(main), " a ")

			}

		}

	}

}
