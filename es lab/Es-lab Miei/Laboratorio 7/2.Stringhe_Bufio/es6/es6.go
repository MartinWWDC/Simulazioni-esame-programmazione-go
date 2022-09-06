package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f := LeggiTesto()
	inf := InvertiStringa(f)
	fmt.Println(inf)
}

func LeggiTesto() (txt string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		} else {

			txt += scanner.Text() + "\n"
		}
	}
	return
}

func InvertiStringa(s string) (txt string) {
	txt = ""
	for i := len(s) - 1; i >= 0; i-- {
		txt += string(s[i])
	}
	return
}
