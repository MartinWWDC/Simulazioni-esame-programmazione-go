package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Istruzione struct {
	iS   string
	par1 string
	par2 string
}
type Stato struct {
	iS        []Istruzione
	ptr       int
	variabili map[string]int
}

func main() {
	Is := LeggiIstruzioni()
	fmt.Println(Is)
}

func LeggiIstruzioni() (istruzioni []Istruzione) {
	Scanner := bufio.NewScanner(os.Stdin)
	for Scanner.Scan() {
		arr := strings.Split(Scanner.Text(), " ")
		if arr[0] == "PRT" {
			arr = append(arr, "")
		}
		istruzioni = append(istruzioni, Istruzione{arr[0], arr[1], arr[2]})
	}
	return
}

func InizializzaStato(programma []Istruzione) (s Stato) {
	variabili := make(map[string]int)
	return Stato{programma, 0, variabili}

}
func Execute(s *Stato) {
	ex := s.iS[s.ptr]
	switch ex.iS {
	case "VAR":
		s.variabili[ex.par1], _ = strconv.Atoi(ex.par2)

	case "ASS":
		s.variabili[ex.par1], _ = s.variabili[ex.par2]

	case "MUL":
		s.variabili[ex.par1] *= s.variabili[ex.par2]

	case "SUB":
		s.variabili[ex.par1] -= s.variabili[ex.par2]
	case "TRI":
		s.variabili[ex.par1] *= 3
	case "PRT":
		fmt.Println(s.variabili[ex.par1])
	case "JMP":
		nomeVar := ex.par1
		numeroIst, _ := strconv.Atoi(ex.par2)
		if nomeVar == "0" {
			if numeroIst > len(s.iS) {
				s.ptr = numeroIst - 1
			}

		}
	}
	if !(s.ptr+1 > len(s.iS)) {

		s.ptr++
	}
}

func Run() {

}
