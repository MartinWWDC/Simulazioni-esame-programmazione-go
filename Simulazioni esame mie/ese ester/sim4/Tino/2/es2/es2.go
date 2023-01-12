package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Comando struct {
	direzione string
	passi     int
}

func main() {
	list := LeggiComandi()
	mapp, min, arrInv := GetStats(list)
	for i, p := range mapp {
		fmt.Println(i, p)
	}
	//fmt.Println(mapp)
	fmt.Println(min)
	for _, i := range arrInv {
		fmt.Print(i)
	}
}

func LeggiComandi() (list []Comando) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		pass, _ := strconv.Atoi(arr[1])
		list = append(list, Comando{arr[0], pass})
	}
	return
}
func GetStats(list []Comando) (map[string]int, Comando, []Comando) {
	mapp := make(map[string]int)
	var min Comando

	for _, i := range list {

		mapp[i.direzione] += i.passi

	}
	g := 0
	for d, p := range mapp {
		if g == 0 {
			min = Comando{d, p}
		} else {
			if p < min.passi {
				min = Comando{d, p}

			}
		}

		g++
	}
	arrInv := []Comando{}
	for i := len(list) - 1; i >= 0; i-- {
		j := list[i].passi
		n := ""
		switch list[i].direzione {
		case "EST":
			n = "OVEST"
		case "NORD":
			n = "SUD"
		case "SUD":
			n = "NORD"
		case "OVEST":
			n = "EST"

		}
		arrInv = append(arrInv, Comando{n, j})
	}
	return mapp, min, arrInv
}

func isIn(n string, arr []string) bool {
	for _, j := range arr {
		if j == n {
			return true
		}
	}
	return false
}
