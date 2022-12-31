package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Punto struct {
	et string
	x  float64
	y  float64
}
type Triangolo struct {
	p1 Punto
	p2 Punto
	p3 Punto
}

func main() {
	//(SortString("ADC"))
	punti := LeggiPunti()
	tr := (GeneraTriangoli(punti))
	//fmt.Println(StringPunto(punti[0]))
	for _, T := range tr {

		if (T.p1.x == T.p2.x || T.p2.x == T.p3.x || T.p3.x == T.p1.x) && (T.p1.y == T.p2.y || T.p2.y == T.p3.y || T.p3.y == T.p1.y) {
			//fmt.Println(T)
			gP1 := getGrade(T.p1)
			gP2 := getGrade(T.p2)
			gP3 := getGrade(T.p3)
			//fmt.Println(gP1, gP2, gP3)
			if gP1 == gP2 && gP2 == gP3 && gP3 == gP1 {
				fmt.Println(StringTriangolo(T))
			}
		}
	}
}

func LeggiPunti() (punti []Punto) {
	Scanner := bufio.NewScanner(os.Stdin)
	for Scanner.Scan() {
		arr := strings.Split(Scanner.Text(), ";")
		x, _ := strconv.ParseFloat(arr[1], 64)
		y, _ := strconv.ParseFloat(arr[2], 64)

		punti = append(punti, Punto{arr[0], x, y})
	}

	return
}
func GeneraTriangoli(punti []Punto) (triangoli []Triangolo) {
	mapp := make(map[string]int)

	for _, i := range punti {
		for _, g := range punti {
			for _, h := range punti {

				if i != g && g != h && i != h {
					str := i.et + g.et + h.et
					str = SortString(str)
					mapp[str]++
					if mapp[str] == 1 {
						triangoli = append(triangoli, Triangolo{i, g, h})
					}
				}
			}
		}
	}
	return
}

func SortString(str string) string {
	arr := strings.Split(string(str), "")
	//fmt.Println(arr)

	(sort.Strings(arr))
	//fmt.Println(arr)
	str = ""
	for _, h := range arr {
		str += h
	}
	return str
}
func Distanza(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
func StringPunto(p Punto) string {
	return p.et + "=(" + fmt.Sprint(p.x) + "," + fmt.Sprint(p.y) + ")"
}
func StringTriangolo(t Triangolo) string {
	return "Triangolo rettangolo con vertici: " + StringPunto(t.p1) + " " + StringPunto(t.p2) + " " + StringPunto(t.p3) + "ed area" + fmt.Sprint(Area(t))
}
func Area(t Triangolo) float64 {
	Per := Distanza(t.p1, t.p2) + Distanza(t.p2, t.p3) + Distanza(t.p3, t.p1)
	return math.Sqrt(Per * (Per - Distanza(t.p1, t.p2)) * (Per - Distanza(t.p2, t.p3)) * (Per - Distanza(t.p3, t.p1)))
}
func getGrade(p Punto) int {
	switch {
	case p.x >= float64(0) && p.y >= float64(0):
		return 1
	case p.x >= float64(0) && p.y < float64(0):
		return 4
	case p.x < float64(0) && p.y >= float64(0):
		return 2
	case p.x < float64(0) && p.y < float64(0):
		return 3
	}
	return 0
}
