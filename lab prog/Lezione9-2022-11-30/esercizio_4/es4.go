package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str := os.Args[1]
	Sottostringhe(str)

}

func Sottostringhe(s string) []string {
	str := strings.Split(s, "")
	for i := range str {
		for h := range str {
			sub := s[i:h]
			if len(sub) > 1 {
				h := 0
				g := 0
				check := true
				for r := 0; r < len(sub)-1; r++ {
					h, _ = strconv.Atoi(string(sub[r]))
					g, _ = strconv.Atoi(string(sub[r+1]))
					if h < g {

					} else {
						check = false
					}

				}
				if check {
					fmt.Println(sub)
				}
			}
		}
	}
	return nil
}
