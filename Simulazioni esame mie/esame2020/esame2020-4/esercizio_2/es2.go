package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(romanoToDecimale("CXLI"))
	//fmt.Println(romanoToDecimale("MCMLXXXIII"))
	//fmt.Println(romanoToDecimale("CCCLXXXIX"))
	//fmt.Println(romanoToDecimale("MMMCMXCIX"))
	//fmt.Println(romanoToDecimale("MDCCLXXXIV"))

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
}
func romanoToDecimale(strP string) int {
	n := 0
	str := strings.Split(strP, "")
	arrInt := []int{}
	for i := 0; i < len(str); i++ {
		arrInt = append(arrInt, converti(str[i]))
	}
	for i := 0; i < len(arrInt); i++ {
		switch {
		case i+1 < len(arrInt) && arrInt[i] == 1 && (arrInt[i+1] == 5 || arrInt[i+1] == 10):
			n += arrInt[i+1] - 1
			i++
		case i+1 < len(arrInt) && arrInt[i] == 10 && (arrInt[i+1] == 50 || arrInt[i+1] == 100):
			n += arrInt[i+1] - 10
			i++
		case i+1 < len(arrInt) && arrInt[i] == 100 && (arrInt[i+1] == 500 || arrInt[i+1] == 1000):
			n += arrInt[i+1] - 100
			i++
		default:
			n += arrInt[i]
		}
	}
	return n
}

func converti(r string) int {
	switch r {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func numeroRomanoBenFormattato(num string) bool {
	arr := strings.Split(num, "")
	s := []int{}
	for g := range arr {
		switch {
		case arr[g] == "I":
			s = append(s, 1)
		case arr[g] == "V":
			s = append(s, 5)
		case arr[g] == "X":
			s = append(s, 10)
		case arr[g] == "L":
			s = append(s, 50)
		case arr[g] == "C":
			s = append(s, 100)
		case arr[g] == "D":
			s = append(s, 500)
		case arr[g] == "M":
			s = append(s, 1000)

		}
	}
	for i := 0; i < len(s); i++ {
		switch {

		case i+3 < len(s) && s[i] == s[i+1] && s[i+1] == s[i+2] && s[i+2] == s[i+3]:
			//fmt.Println("1")
			return false

		case i+2 < len(s) && (s[i+2] == 5 || s[i+2] == 10) && s[i] == 1 && s[i+1] == 1:
			//fmt.Println("2")

			return false

		case i+2 < len(s) && (s[i+2] == 50 || s[i+2] == 100) && s[i] == 10 && s[i+1] == 10:
			//fmt.Println("3")

			return false

		case i+2 < len(s) && (s[i+1] == 500 || s[i+1] == 1000) && s[i] == 100 && s[i+1] == 100:
			//fmt.Println("4")

			return false

		case i < len(s)-1 && s[i] == 1 && (s[i+1] == 5 || s[i+1] == 10):

		case i < len(s)-1 && s[i] == 10 && (s[i+1] == 50 || s[i+1] == 100):

		case i < len(s)-1 && s[i] == 100 && (s[i+1] == 500 || s[i+1] == 1000):

		case i < len(s)-1 && s[i] < s[i+1]:
			return false
		}

	}
	return true

}
