package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	g := LeggiTesto()
	t := FiltraTesto(g)
	fmt.Println(t)
	fmt.Println(len(t))

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

func ContieneCifre(s string) bool {
	for _, g := range s {
		if strings.ContainsRune("0123456789", g) {
			return true
		}
	}
	return false
}
func FiltraTesto(sl []string) (arr []string) {
	for _, g := range sl {
		if ContieneCifre(g) {
			arr = append(arr, g)
		}
	}
	return
}
