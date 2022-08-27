package main

import "fmt"

func main() {
	var n string
	vmaiu, vmin, cmaiu, cmin := 0, 0, 0, 0

	fmt.Scan(&n)
	for _, l := range n {
		if èLetteraValida(l) {
			if Vocale(l) {
				if èMaiuscola(l) {
					vmaiu++
				} else {
					vmin++
				}
			} else {
				if èMaiuscola(l) {
					cmaiu++
				} else {
					cmin++
				}
			}
		}
	}
	fmt.Print("Vocali maiuscole: ")
	fmt.Println(vmaiu)
	fmt.Print("Consonanti maiuscole:")
	fmt.Println(cmaiu)
	fmt.Print("Vocali minuscole: ")
	fmt.Println(vmin)
	fmt.Print("Consonanti minuscole: ")
	fmt.Println(cmin)

}

func èLetteraValida(l rune) bool {
	if l >= 'a' && l <= 'z' || l >= 'A' && l <= 'Z' {
		return true
	} else {
		return false
	}
}

func èMaiuscola(l rune) bool {
	if l >= 'A' && l <= 'Z' {
		return true
	} else {
		return false
	}

}

func Vocale(l rune) bool {
	if èMaiuscola(l) {
		l += 32
	}
	switch l {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
