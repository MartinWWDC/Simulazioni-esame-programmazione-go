/**
 * non so come finire
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Utenza struct {
	NUMERO_TELEFONO string
	CODICE_SIM      string
}
type RegistroTelefonico struct {
	Utenze []Utenza
}

func main() {
	utenze := LeggiUtenze()
	r := InizializzaRegistroC(utenze)
	n := 0.0
	search := ""
	maxChanges := 0
	mapp := make(map[string][]string)
	for _, h := range r.Utenze {
		if h.NUMERO_TELEFONO[:len(search)] == search {
			//fmt.Println("s:", len(simsList(r, h.NUMERO_TELEFONO)))
			if len(simsList(r, h.NUMERO_TELEFONO)) > maxChanges {
				maxChanges = len(simsList(r, h.NUMERO_TELEFONO)) - 1

			}
			mapp[h.NUMERO_TELEFONO] = append(mapp[h.NUMERO_TELEFONO], h.CODICE_SIM)
		}
	}
	for _, k := range mapp {
		if n > 1 {
			n += float64(len(k))
		}
	}
	fmt.Println(float64(NumeroUtenze(r, search)))
	fmt.Println(n)
	fmt.Println(maxChanges)
	fmt.Println(len(mapp))
	/*for j, l := range mapp {
		fmt.Println("numero:", j, "SIMS:", len(l))
	}*/
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
	return
}
func InizializzaRegistroC(utenze []Utenza) (registro RegistroTelefonico) {
	registro.Utenze = utenze
	return
}
func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) (registroAggiornato RegistroTelefonico) {
	registroAggiornato = registro
	registroAggiornato.Utenze = append(registroAggiornato.Utenze, utenza)
	return
}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {
	codiceSIM = ""
	for _, h := range registro.Utenze {
		if h.NUMERO_TELEFONO == telefono {
			codiceSIM = h.CODICE_SIM
		}
	}
	return
}

func NumeroUtenze(registro RegistroTelefonico, telefono string) (n int) {
	for _, h := range registro.Utenze {
		if h.NUMERO_TELEFONO[:len(telefono)] == telefono {
			n++
		}
	}
	return
}
func simsList(registro RegistroTelefonico, telefono string) (sims []string) {
	for _, h := range registro.Utenze {
		if h.NUMERO_TELEFONO == telefono {
			sims = append(sims, h.CODICE_SIM)
		}
	}
	return
}
