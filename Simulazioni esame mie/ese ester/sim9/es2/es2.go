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
	g := LeggiComandi()
	fmt.Println(g)
	h := AnalizzaComandi(g)
	fmt.Println(h)

}

func LeggiComandi() (commands []Comando) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		pass, _ := strconv.Atoi(arr[1])
		commands = append(commands, Comando{arr[0], pass})
	}
	return
}

func AnalizzaComandi(comandi []Comando) map[string]int {
	mappa := make(map[string]int)
	for _, g := range comandi {
		mappa[g.direzione] += g.passi
	}
	return mappa
}
