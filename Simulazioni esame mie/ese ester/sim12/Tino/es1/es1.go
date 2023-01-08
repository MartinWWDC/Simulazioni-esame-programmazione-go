package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arr := os.Args[1:]
	for h := range arr {
		fmt.Print(TrasformaParola(arr[h], h), " ")

	}
}

func TrasformaParola(parola string, posizione int) (parolaTraformata string) {
	for j := range parola {
		if posizione%2 == 0 {

			if j%2 == 0 {
				parolaTraformata += strings.ToUpper(string(parola[j]))
			} else {
				parolaTraformata += strings.ToLower(string(parola[j]))

			}
		} else {
			if j%2 == 0 {
				parolaTraformata += strings.ToLower(string(parola[j]))
			} else {
				parolaTraformata += strings.ToUpper(string(parola[j]))

			}
		}
	}

	return
}
