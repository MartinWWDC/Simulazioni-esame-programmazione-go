package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Utenza struct {
	telefono, sim string
}

type RegistroTelefonico map[string][]string

func main() {
	uten := LeggiUtenze()
	reg := InizializzaRegistro()
	for i := 0; i < len(uten); i++ {
		reg = AggiungiUtenza(reg, uten[i])
	}
	//fmt.Println(uten)

	for i := 0; i < len(uten); i++ {
		si := uten[i].telefono[0:3]
		if si == "338" {
			numero := uten[i].telefono
			fmt.Println("Il numero", numero, " è associato alla sim", Identifica(reg, numero))
		}
	}

}
func LeggiUtenze() (utenze []Utenza) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ";")
		utenze = append(utenze, Utenza{text[0], text[1]})
	}
	return utenze
}

func InizializzaRegistro() (registro RegistroTelefonico) {
	registro = make(RegistroTelefonico)
	return
}

func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) RegistroTelefonico {
	/*
		AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) (registroAggiornato RegistroTelefonico) che riceve in input
		un'istanza di tipo RegistroTelefonico nella variabile registro e un'istanza di tipo Utenza nella variabile utenza.
		Se il numero dell'utenza non è presente nel registro, viene aggiunta una nuova voce.
		Altrimenti, viene solamente aggiunta la nuova sim al numero di telefono.
		La funzione restituisce il registro telefonico aggiornato nella variabile registroAggiornato.
	*/
	registro[utenza.telefono] = append(registro[utenza.telefono], utenza.sim)

	return registro
}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {

	/*
	   Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) che riceve in input un'istanza di tipo RegistroTelefonico
	   nella variabile registro e un valore di tipo string nella variabile telefono (un numero di telefono mobile).
	   La funzione restituisce il codice della SIM (presente in registro) corrispondente a telefono nella variabile codiceSIM.
	   Se in registro sono presenti più codici SIM corrispondenti a telefono, viene restituito il codice più recente, ovvero quello
	   che nella sequenza di input è letto per ultimo.
	   Se in registro non è presente alcun codice SIM corrispondente a telefono, viene restituito il valore "".
	*/

	sim := (registro[telefono])
	codiceSIM = (sim[len(sim)-1])
	return
}
