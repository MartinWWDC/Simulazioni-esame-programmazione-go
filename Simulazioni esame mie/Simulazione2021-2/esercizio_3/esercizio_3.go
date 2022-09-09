package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Utenza struct {
	num string
	cod string
}
type RegistroTelefonico map[string][]string

func main() {
	utenze := LeggiUtenze()
	registro := InizializzaRegistro()

	for _, utenza := range utenze {
		AggiungiUtenza(registro, utenza)
	}

	for telefono, _ := range registro {
		if telefono[:3] == "338" {
			simRecente := Identifica(registro, telefono)
			fmt.Println("Il numero", telefono, "è associato alla sim", simRecente)
		}
	}
}
func LeggiUtenze() (utenze []Utenza) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		utenze = append(utenze, Utenza{arr[0], arr[1]})
	}
	return

}
func InizializzaRegistro() (registro RegistroTelefonico) {
	registro = make(RegistroTelefonico)
	return
}
func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) RegistroTelefonico {
	registro[utenza.num] = append(registro[utenza.num], utenza.cod)
	return registro
}
func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {
	codiceSIM = registro[telefono][len(registro[telefono])-1]
	return

}
