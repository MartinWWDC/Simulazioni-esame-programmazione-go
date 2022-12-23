package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	str := LeggiTesto()
	words := SeparaParole(str)
	print(occorrenze(words))
}
func LeggiTesto() string {
	Scanner := bufio.NewScanner(os.Stdin)
	str := ""
	for Scanner.Scan() {
		str += Scanner.Text() + " "
	}
	return str
}
func SeparaParole(s string) (words []string) {
	arr := strings.Split(s, " ")
	c := true
	var j int
	for _, h := range arr {
		if h != "" {
			for n, i := range h {
				if !unicode.IsLetter(i) && i != '!' && i != '?' && i != ',' {
					c = false
					break
				} else {
					if i == '!' || i == '?' || i == ',' {
						j = n
					}
				}
			}
			if c {
				if j != 0 {
					words = append(words, h[:j])
					j = 0
				} else {
					words = append(words, h)

				}
			}
		}
		c = true

	}

	return words
}
func occorrenze(words []string) map[string]int {
	mapp := make(map[string]int)
	for _, i := range words {
		mapp[i]++
	}
	return mapp
}
func print(mapp map[string]int) {
	for h, i := range mapp {
		fmt.Print(string(h), " :")
		for x := 0; x < i; x++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
