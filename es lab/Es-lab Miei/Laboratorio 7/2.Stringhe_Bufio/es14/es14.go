package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	date := LeggiTesto()

	var dateFormattate []string
	for _, data := range date {
		dateFormattate = append(dateFormattate, Formatta(data))
	}

	sort.Strings(dateFormattate)
	for _, data := range dateFormattate {
		fmt.Println(data)
	}
}

func LeggiTesto() (date []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			return
		}
		date = append(date, riga)
	}
	return
}

func Formatta(data string) string {
	var anno, mese, giorno string
	dataComposta := strings.Split(data, "/")
	fmt.Println(dataComposta)

	if len(dataComposta[0]) <= 2 {
		anno = dataComposta[2]
		mese = dataComposta[1]
		giorno = dataComposta[0]

	} else if len(dataComposta[0]) == 4 {
		anno = dataComposta[0]
		mese = dataComposta[1]
		giorno = dataComposta[2]
	}
	if len(mese) < 2 {
		mese = "0" + mese
	}

	if len(giorno) < 2 {
		giorno = "0" + giorno
	}
	return anno + "-" + mese + "-" + giorno
}
