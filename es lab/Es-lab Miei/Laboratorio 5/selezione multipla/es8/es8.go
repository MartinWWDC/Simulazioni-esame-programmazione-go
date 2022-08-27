package main

import "fmt"

func main() {
	var n string
	fmt.Print("Inserire mese: ")
	fmt.Scan(&n)
	fmt.Println(numeroDiGiorni(n))
}

func numeroDiGiorni(mese string) int {
	if mese == "gennaio" || mese == "febbraio" || mese == "marzo" || mese == "aprile" || mese == "maggio" || mese == "giugno" || mese == "luglio" || mese == "agosto" || mese == "settembre" || mese == "ottobre" || mese == "novembre" || mese == "dicembre" {

		switch mese {
		case "settembre", "aprile", "giugno", "novembre":
			return 30
		case "febbraio":
			return 28
		default:
			return 31
		}

	} else {
		return 0
	}
}
