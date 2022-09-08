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
			if i < len(iarr)-1 {
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
	check := true
	var iarr []int
	arr := strings.Split(s, "")
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
	//var check
	//for i := len(iarr) - 1; i > 0; i-- {
	//
	//	default:
	//		result += iarr[i]
	//
	//	}
	//}

	for i := 0; i < len(iarr)-1; i++ {
		switch {
		case i < len(iarr)-3 && iarr[i] == iarr[i+1] && iarr[i+1] == iarr[i+2] && iarr[i+2] == iarr[i+3]:
			return false
		case i < len(iarr)-2 && iarr[i] == 1 && iarr[i+1] == 1 && (iarr[i+2] == 5 || iarr[i+2] == 10): // IIV o IIX
			return false
		case i < len(iarr)-2 && iarr[i] == 10 && iarr[i+1] == 10 && (iarr[i+2] == 50 || iarr[i+2] == 100): // XXL o XXC
			return false
		case i < len(iarr)-2 && iarr[i] == 100 && iarr[i+1] == 100 && (iarr[i+2] == 500 || iarr[i+2] == 1000): // CCD o CCM
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
	/*

		for i := 0; i < len(iarr); i++ {

			switch {
			case i < len(iarr)-1 && iarr[i] == 1 && (iarr[i+1] == 5 || iarr[i+1] == 10):
				check = true
			case i < len(iarr)-1 && iarr[i] == 10 && (iarr[i+1] == 50 || iarr[i+1] == 100):
				check = true
			case i < len(iarr)-1 && iarr[i] == 100 && (iarr[i+1] == 500 || iarr[i+1] == 1000):
				check = true
			case i > 0 && (iarr[i] == 5 || iarr[i] == 50 || iarr[i] == 500) && iarr[i-1] < iarr[i]:
				check = true
			default:

			}

		}*/
}
