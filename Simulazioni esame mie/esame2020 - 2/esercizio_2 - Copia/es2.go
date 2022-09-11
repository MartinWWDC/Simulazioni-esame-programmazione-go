package main

import "fmt"

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

func Convert(str string) (att []int) {
	for _, g := range str {
		switch {
		case g == 'I':
			att = append(att, 1)
		case g == 'V':
			att = append(att, 5)

		case g == 'X':
			att = append(att, 10)

		case g == 'L':
			att = append(att, 50)

		case g == 'C':
			att = append(att, 100)

		case g == 'D':
			att = append(att, 500)

		case g == 'M':
			att = append(att, 1000)

		}
	}
	return
}

func romano2decimale(gin string) (result int) {
	iarr := Convert(gin)
	//fmt.Println(iarr)
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
	return result
}

func numeroRomanoBenFormattato(gin string) bool {
	iarr := Convert(gin)
	//fmt.Println(iarr)
	check := true
	for i := 0; i < len(iarr); i++ {

		switch {

		case i < len(iarr)-1 && iarr[i] == 1 && (iarr[i+1] == 5 || iarr[i+1] == 10):

		case i < len(iarr)-1 && iarr[i] == 10 && (iarr[i+1] == 50 || iarr[i+1] == 100):

		case i < len(iarr)-1 && iarr[i] == 100 && (iarr[i+1] == 500 || iarr[i+1] == 1000):

		case i < len(iarr)-2 && iarr[i] == 1 && iarr[i] == iarr[i+1] && iarr[i+1] == iarr[i+2] && (iarr[i+1] == 500 || iarr[i+1] == 1000):
			check = false

		case i < len(iarr)-2 && iarr[i] == 10 && iarr[i] == iarr[i+1] && iarr[i+1] == iarr[i+2] && (iarr[i+1] == 500 || iarr[i+1] == 1000):
			check = false

		case i < len(iarr)-2 && iarr[i] == 100 && iarr[i] == iarr[i+1] && iarr[i+1] == iarr[i+2] && (iarr[i+1] == 500 || iarr[i+1] == 1000):
			check = false
		case i < len(iarr)-3 && iarr[i] == iarr[i+1] && iarr[i+1] == iarr[i+2] && iarr[i+2] == iarr[i+3]:
			check = false
		case i < len(iarr)-1 && iarr[i] < iarr[i+1]:
			return false

		default:

		}
	}
	return check
}
