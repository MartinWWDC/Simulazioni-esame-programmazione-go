package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	et string
	x  float64
	y  float64
}

func main() {
	perc := leggiInput()
	fmt.Println(perc)
	dist := Lunghezza(perc)
	fmt.Println("distanza: ", dist)
	var sum float64
	for g := range perc {
		if sum >= dist/2 {
			fmt.Println("sex:", toString(perc[g]))
			break
		} else {
			if g != len(perc)-1 {
				sum += (Distanza(perc[g], perc[g+1]))

			}
		}

	}
}

func leggiInput() (percorso []Punto) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		x, _ := strconv.ParseFloat(arr[1], 64)
		y, _ := strconv.ParseFloat(arr[2], 64)
		percorso = append(percorso, Punto{arr[0], x, y})
	}
	return
}

func Lunghezza(tragitto []Punto) (sum float64) {
	sum = 0.0
	for g := range tragitto {
		if g != len(tragitto)-1 {
			sum += (Distanza(tragitto[g], tragitto[g+1]))

		}
	}
	sum = sum * 1000
	sum = math.Round(sum)
	sum /= 1000
	return
}

func Distanza(p1, p2 Punto) float64 {
	x := math.Pow((p1.x - p2.x), 2)
	y := math.Pow((p1.y - p2.y), 2)
	//fmt.Println(x, y)
	return math.Sqrt(y + x)
}
func toString(p Punto) string {
	x := fmt.Sprint(p.x)
	y := fmt.Sprint(p.y)
	return p.et + " = (" + x + " , " + y + ")"
}
