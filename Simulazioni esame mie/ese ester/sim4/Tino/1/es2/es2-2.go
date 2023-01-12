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
	scanner := bufio.NewScanner(os.Stdin)
	mapp := make(map[string]int)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(arr[1])
		mapp[arr[0]] += n
	}
	fmt.Println(mapp)
}

func LeggiComandi() (comandi []Comando) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(arr[1])
		comandi = append(comandi, Comando{arr[0], n})
	}
	return
}
func AnalizzaComandi(comandi []Comando) map[string]int {
	mapp := make(map[string]int)
	for _, h := range comandi {
		mapp[h.direzione] += h.passi

	}
	return mapp
}
func getStats(mapp map[string]int) {
	var lower Comando
	for h := range comandi {
		if h == 0 {
			lower=
		}
	}
}
