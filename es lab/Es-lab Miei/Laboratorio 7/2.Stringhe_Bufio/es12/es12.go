package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//g := LeggiTesto()
	//fmt.Println(g)
	s := Formatta("2019/04/3")
	fmt.Println(s)
}
func LeggiTesto() (arr []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			arr = append(arr, scanner.Text())
		} else {
			break
		}
	}
	return
}

func Formatta(s string) string {
	var sform []rune
	run := []rune(s)
	if len(s) == 8 {
		sform = append(run[:4], '-', '0', run[5], '-', '0', run[7])
	}
	if len(s) == 10 {

		sform = append(run[:4], '-', run[5], run[6], '-', run[8], run[9])

	}

	if len(s) == 9 {
		if run[6] == '/' {
			sform = append(run[:4], '-', '0', run[5], '-', run[7], run[8])

		} else {
			sform = append(run[:4], '-', run[5], run[6], '-', '0', run[8])

		}
	}

	return string(sform)
}
