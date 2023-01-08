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
	tr := NuovoTragitto()
	//fmt.Println(tr)
	fmt.Println(Lunghezza(tr))
	middle := (getMiddle(tr))
	fmt.Println(String(middle))
}

func NuovoTragitto() (tragitto []Punto) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		x, _ := strconv.ParseFloat(arr[1], 64)
		y, _ := strconv.ParseFloat(arr[2], 64)
		tragitto = append(tragitto, Punto{arr[0], x, y})
	}
	return
}

func Distanza(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func String(p Punto) string {
	return p.et + "= (" + fmt.Sprint(p.x) + "," + fmt.Sprint(p.y) + ")"
}

func Lunghezza(tragitto []Punto) float64 {
	dist := 0.0
	for i := 1; i <= len(tragitto)-1; i++ {
		//	fmt.Println(tragitto[i].et)
		dist += Distanza(tragitto[i-1], tragitto[i])
	}
	return dist
}
func getMiddle(tragitto []Punto) Punto {
	dist := Lunghezza(tragitto)
	middle := 0.0
	for i := 1; i <= len(tragitto)-1; i++ {
		//fmt.Println(tragitto[i].et)
		middle += Distanza(tragitto[i-1], tragitto[i])
		if middle > dist/2 {
			return tragitto[i]
		}
	}
	return Punto{}
}
