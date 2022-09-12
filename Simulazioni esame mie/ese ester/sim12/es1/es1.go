package main

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:len(os.Args)]
	fmt.Println((arr))
	for i, g := range arr {
		tr := TrasformaParola(g, i)
		fmt.Print(tr, " ")
	}

}
func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
	for i, g := range parola {
		if posizione%2 != 0 {
			if i%2 != 0 {
				parolaTrasformata += string(g - 32)
			} else {
				parolaTrasformata += string(g)
			}
		} else {
			if i%2 == 0 {
				parolaTrasformata += string(g - 32)
			} else {
				parolaTrasformata += string(g)
			}
		}
		//fmt.Println(parolaTrasformata)
	}
	return
}
