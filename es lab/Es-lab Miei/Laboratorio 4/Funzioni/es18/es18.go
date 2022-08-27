package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci la soglia: ")
	fmt.Scan(&n)
	TernePitagoriche(n)

}
func ÈTernaPitagoriga(a int, b int, c int) bool {
	if (c * c) == a*a+b*b {
		return true
	} else {
		return false
	}
}

func TernePitagoriche(soglia int) {
	for i := 1; i < soglia; i++ {
		for j := 1; j < soglia; j++ {
			for z := 1; z < soglia; z++ {
				if ÈTernaPitagoriga(i, j, z) {
					fmt.Print("(")
					fmt.Print(i)
					fmt.Print(",")
					fmt.Print(j)
					fmt.Print(",")
					fmt.Print(z)
					fmt.Println(")")

				}
			}
		}
	}
}
