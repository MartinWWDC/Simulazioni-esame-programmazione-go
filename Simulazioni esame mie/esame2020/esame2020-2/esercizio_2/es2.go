package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("CXLI:", numeroRomanoBenFormattato("CXLI"))
	fmt.Println("MCMLXXXIII:", numeroRomanoBenFormattato("MCMLXXXIII"))
	fmt.Println("CCCLXXXIX:", numeroRomanoBenFormattato("CCCLXXXIX"))
	fmt.Println("MMMCMXCIX:", numeroRomanoBenFormattato("MMMCMXCIX"))
	fmt.Println("MDCCLXXXIV:", numeroRomanoBenFormattato("MDCCLXXXIV"))
	fmt.Println("MIIII:", numeroRomanoBenFormattato("MIIII"))
	fmt.Println("VXII:", numeroRomanoBenFormattato("VXII"))
	fmt.Println("MIC:", numeroRomanoBenFormattato("MIC"))
	fmt.Println("MDICCLXXXIX:", numeroRomanoBenFormattato("MDICCLXXXIX"))
	fmt.Println("CLDXIX:", numeroRomanoBenFormattato("CLDXIX"))

}

func romano2decimale(n string) int {
	var iarr []int
	result := 0
	arr := strings.Split(n, "")
	for g := range arr {
		switch {
		case arr[g] == "I":
			iarr = append(iarr, 1)
		case arr[g] == "V":
			iarr = append(iarr, 5)
		case arr[g] == "X":
			iarr = append(iarr, 10)
		case arr[g] == "L":
			iarr = append(iarr, 50)
		case arr[g] == "C":
			iarr = append(iarr, 100)
		case arr[g] == "D":
			iarr = append(iarr, 500)
		case arr[g] == "M":
			iarr = append(iarr, 1000)

		}
	}
	for i := 0; i < len(iarr); i++ {
		fmt.Println(result)

		switch {
		case i < len(iarr)-1 && iarr[i] == 1 && (iarr[i+1] == 5 || iarr[i+1] == 10):
			result += iarr[i+1] - 1
			i++

		case i < len(iarr)-1 && iarr[i] == 10 && (iarr[i+1] == 50 || iarr[i+1] == 100):
			result += iarr[i+1] - 10
			if i < len(iarr)-1 {
				i++

			}
		case i < len(iarr)-1 && iarr[i] == 100 && (iarr[i+1] == 500 || iarr[i+1] == 1000):
			result += iarr[i+1] - 100
			if i <3 len(iarr)-1 {
				i++

			}
		default:
			result += iarr[i]

		}

	}

	fmt.Println(iarr)
	return result
}

func numeroRomanoBenFormattato(s string) bool {
	var prev byte
	rip := 0
	check := true
	
		for g := range arr {
			switch {
			case arr[g] == "I":
				iarr = append(iarr, 1)
			case arr[g] == "V":
				iarr = append(iarr, 5)
			case arr[g] == "X":
				iarr = append(iarr, 10)
			case arr[g] == "L":
				iarr = append(iarr, 50)
			case arr[g] == "C":
				iarr = append(iarr, 100)
			case arr[g] == "D":
				iarr = append(iarr, 500)
			case arr[g] == "M":
				iarr = append(iarr, 1000)

			}
		}



			check := true
			for i := 0; i < len(iarr)-1; i++ {
				switch {
				case i < len(iarr)-3 && iarr[i] == iarr[i+1] && iarr[i+1] == iarr[i+2] && iarr[i+2] == iarr[i+3]: // controllo che un elemento non sia ripetuto per + di 3 volte
					return false
				case i < len(iarr)-2 && iarr[i] == 1 && iarr[i+1] == 1 && (iarr[i+2] == 5 || iarr[i+2] == 10): // IIV o IIX controllo che non ci siamo 2 volte 1  prima di 5 e 10
					return false
				case i < len(iarr)-2 && iarr[i] == 10 && iarr[i+1] == 10 && (iarr[i+2] == 50 || iarr[i+2] == 100): // XXL o XXC controllo che non ci siamo 2 volte 10  prima di 50 e 100
					return false
				case i < len(iarr)-2 && iarr[i] == 100 && iarr[i+1] == 100 && (iarr[i+2] == 500 || iarr[i+2] == 1000): // CCD o CCM controllo che non ci siamo 2 volte 100  prima di 500 e 1000
					return false
				case i < len(iarr)-1 && iarr[i] == 1 && (iarr[i+1] == 5 || iarr[i+1] == 10):

				case i < len(iarr)-1 && iarr[i] == 10 && (iarr[i+1] == 50 || iarr[i+1] == 100):

				case i < len(iarr)-1 && iarr[i] == 100 && (iarr[i+1] == 500 || iarr[i+1] == 1000):

				case iarr[i] < iarr[i+1]:
					return false
				default:

				}
			}
			return check
}
