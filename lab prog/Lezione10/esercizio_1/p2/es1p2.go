package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func main() {
	str := LeggiTesto()
	mapp := Occoreza(str)
	print(mapp)
}
func LeggiTesto() string {
	Scanner := bufio.NewScanner(os.Stdin)
	str := ""
	for Scanner.Scan() {
		str += Scanner.Text()

	}
	return str
}

func Occoreza(s string) map[rune]int {
	mapp := make(map[rune]int)
	//str := strings.Split(s, "")
	for _, i := range s {
		if unicode.IsLetter(i) {
			mapp[i]++
		}
	}
	return mapp
}

func print(mapp map[rune]int) {
	keys := []string{}
	for h, _ := range mapp {
		keys = append(keys, string(h))

	}
	sort.Strings(keys)
	for _, h := range keys {
		r := []rune(h)
		fmt.Print(string(h), " :")
		for x := 0; x < mapp[r[0]]; x++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
