package main

import (
	"fmt"
	"os"
)

func main() {
	/*
			fmt.Println(romano2Decimale("CXLI") == 141)
			fmt.Println(romano2Decimale("MCMLXXXIII") == 1983)
			fmt.Println(romano2Decimale("CCCLXXXIX") == 389)
			fmt.Println(romano2Decimale("MMMCMXCIX") == 3999)
			fmt.Println(romano2Decimale("MDCCLXXXIV") == 1784)

		fmt.Println(numeroRomanoBenFormattato("CXLI"))
		fmt.Println(numeroRomanoBenFormattato("MCMLXXXIII"))
		fmt.Println(numeroRomanoBenFormattato("CCCLXXXIX"))
		fmt.Println(numeroRomanoBenFormattato("MMMCMXCIX"))
		fmt.Println(numeroRomanoBenFormattato("MDCCLXXXIV"))
		fmt.Println(numeroRomanoBenFormattato("MIIII"))
		fmt.Println(numeroRomanoBenFormattato("VXII"))
		fmt.Println(numeroRomanoBenFormattato("MIC"))
		fmt.Println(numeroRomanoBenFormattato("MDICCLXXXIX"))
		fmt.Println(numeroRomanoBenFormattato("CLDXIX"))
	*/
	n := os.Args[1]
	if numeroRomanoBenFormattato(n) {
		fmt.Println(romano2Decimale(n))
	}

}
func romano2Decimale(n string) int {
	num := 0
	for i := 0; i <= len(n)-1; i++ {
		switch {
		case i < len(n)-1 && convert(rune(n[i])) == 1 && (convert(rune(n[i+1])) == 5 || convert(rune(n[i+1])) == 10):
			num += convert(rune(n[i+1])) - convert(rune(n[i]))
			i++
		case i < len(n)-1 && convert(rune(n[i])) == 10 && (convert(rune(n[i+1])) == 50 || convert(rune(n[i+1])) == 100):
			num += convert(rune(n[i+1])) - convert(rune(n[i]))
			i++
		case i < len(n)-1 && convert(rune(n[i])) == 100 && (convert(rune(n[i+1])) == 500 || convert(rune(n[i+1])) == 1000):
			num += convert(rune(n[i+1])) - convert(rune(n[i]))
			i++
		default:
			num += convert(rune(n[i]))
		}
	}
	return num
}
func convert(s rune) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
func numeroRomanoBenFormattato(n string) bool {
	for i := 0; i <= len(n)-1; i++ {
		switch {
		case i < len(n)-1 && convert(rune(n[i])) == 1 && (convert(rune(n[i+1])) == 5 || convert(rune(n[i+1])) == 10):
		case i < len(n)-1 && convert(rune(n[i])) == 10 && (convert(rune(n[i+1])) == 50 || convert(rune(n[i+1])) == 100):
		case i < len(n)-1 && convert(rune(n[i])) == 100 && (convert(rune(n[i+1])) == 500 || convert(rune(n[i+1])) == 1000):
		case i < len(n)-1 && convert(rune(n[i])) < convert(rune(n[i+1])):
			return false
		case i < len(n)-3 && convert(rune(n[i])) == convert(rune(n[i+1])) && convert(rune(n[i+1])) == convert(rune(n[i+2])) && convert(rune(n[i+2])) == convert(rune(n[i+3])):
			return false

		}
	}
	return true
}
