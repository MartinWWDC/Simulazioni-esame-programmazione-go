package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Istruzione struct {
	istr  string
	param []string
}

type Stato struct {
	istruzioni []Istruzione
	ptr        int
	variabili  map[string]int
}

func LeggiIstruzioni() (istruzioni []Istruzione) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		istruzioni = append(istruzioni, Istruzione{s[0], s[1:]})
	}
	return
}

func InizializzaStato(programma []Istruzione) (s Stato) {
	s.istruzioni = programma
	s.ptr = 0
	s.variabili = make(map[string]int)
	return
}

func EseguiIstruzione(s Stato) (Stato, string) {
	is := s.istruzioni[s.ptr]
	rer := ""
	switch is.istr {
	case "VAR":
		if _, err := s.variabili[is.param[0]]; err {
			rer = "ERRORE - Istruzione" + strconv.Itoa(s.ptr) + ": Ridefinizione della variabile" + is.param[0]
			break
		}
		s.variabili[is.param[0]], _ = strconv.Atoi(is.param[1])
	case "ASS":
		s.variabili[is.param[1]] = s.variabili[is.param[0]]
	case "MUL":
		s.variabili[is.param[0]] *= s.variabili[is.param[1]]
	case "SUB":
		s.variabili[is.param[0]] -= s.variabili[is.param[1]]
	case "TRI":
		s.variabili[is.param[0]] *= 3

	case "PRT":
		fmt.Println(s.variabili[is.param[0]])
	case "JMP":
		if s.variabili[is.param[0]] == 0 {
			s.ptr, _ = strconv.Atoi(is.param[1])

			if len(s.istruzioni)-1 > s.ptr {
				rer = "ERRORE - Istruzione" + strconv.Itoa(s.ptr) + " Riferimento ad un'istruzione inesistente " + is.istr + is.param[0] + is.param[1]
			}
			s.ptr--

		}
	}
	s.ptr++
	//StampS(s)

	return s, rer
}

func EseguiInterprete(programma []Istruzione) {
	s := InizializzaStato(programma)
	var err string
	n, _ := strconv.Atoi(os.Args[1])
	//for {
	for i := 0; i <= n; i++ {
		s, err = EseguiIstruzione(s)
		fmt.Println(err)
		StampS(s)
	}

	/*if err != "" {
		fmt.Println(err)
		break
	}*/
	//}

	StampS(s)
}

func StampS(s Stato) {
	fmt.Println("Istruzioni:", s.istruzioni)
	fmt.Println("ptr:", s.ptr)
	fmt.Println("variabili:", s.variabili)

}
func main() {
	is := LeggiIstruzioni()
	EseguiInterprete(is)
}
