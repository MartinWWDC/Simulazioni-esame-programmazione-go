package main

import "fmt"

func main() {
	var n string

	fmt.Scan(&n)
	fmt.Println(ÈPalindroma(n))

}

/*
func ÈPalindroma(s string) bool {
	//fmt.Print(string([]rune(s)[1]))
	var n string = ""
	var l string
	for i := len(s) - 1; i >= 0; i-- {
		//fmt.Println(i)
		l = (string([]rune(s)[i]))
		n += l
	}
	//fmt.Print(n)
	if n == s {
		return true

	} else {
		return false
	}
}
*/
func ÈPalindroma(parola string) bool {
	parolaCopia := ""
	for _, lettera := range parola {
		parolaCopia = string(lettera) + parolaCopia
	}
	return parola == parolaCopia

}
