package main

import "fmt"

func main() {
	var c rune = 127153
	for i := 0; i < 10; i++ {
		fmt.Print("Simbolo: ")
		fmt.Print(string(c))
		fmt.Print("Codice numerico in base 10: ")
		fmt.Println(c)
		c++
	}
}
