package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Utenza struct {
	num     string
	cod_sim string
}
type RegistroTelefonico map[string][]string

func main() {
	utenze := LeggiUtenze()
	registro := InizializzaRegistro()
	for _, i := range utenze {
		registro = AggiungiUtenza(registro, i)
	}

	ut := getSim(registro, "338")
	fmt.Println(ut)
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
	return make(map[string][]string)
}
func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) (registroAggiornato RegistroTelefonico) {
	registroAggiornato = registro
	registroAggiornato[utenza.num] = append(registro[utenza.num], utenza.cod_sim)
	return

}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {
	return registro[telefono][len(registro[telefono])-1]
}

func getSim(registro RegistroTelefonico, num string) (ut []Utenza) {
	h := 0

	for i := range registro {
		if i[:len(num)] == num {
			ut = append(ut, Utenza{i, registro[i][len(registro[i])-1]})
		}
	}
	fmt.Println(h)
	return ut
}
