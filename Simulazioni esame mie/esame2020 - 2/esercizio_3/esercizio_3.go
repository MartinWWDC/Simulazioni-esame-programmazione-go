package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Posizione struct {
	etichetta   string
	latitudine  float64
	longitudine float64
}
type percorso []Posizione

func main() {
	var arr percorso
	var arr2 percorso
	dislat, _ := strconv.ParseFloat(os.Args[1], 64)
	dislog, _ := strconv.ParseFloat(os.Args[2], 64)
	//fmt.Println(dislat)
	//fmt.Println(dislog)
	pos := scan()
	//fmt.Println(pos)

	for i := 0; i < len(pos); i++ {
		arr = append(arr, pos[i])
		//fmt.Println(pos[i])
		for j := i + 1; j < len(pos)-1; j++ {
			latdif := pos[j].latitudine - pos[i].latitudine
			logdif := pos[j].longitudine - pos[i].longitudine
			//fmt.Println("-", dislat, "< latdif", pos[j].latitudine, "-", pos[i].latitudine, "::", latdif, "<", dislat)
			//fmt.Println("-", dislog, "< dif long", pos[j].longitudine, "-", pos[i].longitudine, "::", logdif, "<", dislog)
			//minusdislat := -dislat
			//fmt.Println(minusdislat)

			if -dislat <= latdif && latdif <= dislat && -dislog <= logdif && logdif <= dislog {
				//fmt.Println("add", pos[j])
				arr = append(arr, pos[j])
			}
		}
		//fmt.Println("arr:", arr)
		fmt.Println(StringPercorso(arr))
		arr = arr2
	}

}
func scan() percorso {
	scanner := bufio.NewScanner(os.Stdin)
	var posArr percorso
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ";")
		lat, _ := strconv.ParseFloat(text[1], 64)
		long, _ := strconv.ParseFloat(text[2], 64)
		pos, check := NuovaPosizione(text[0], lat, long)
		if check {
			posArr = append(posArr, pos)

		}

	}
	return posArr
}
func NuovaPosizione(etichetta string, latitudine float64, longitudine float64) (Posizione, bool) {
	check := false
	if latitudine > -90 && latitudine < 90 && longitudine > -180 && longitudine < 180 {
		check = true
	}
	pos := Posizione{etichetta, latitudine, longitudine}
	return pos, check
}

func StringPosizione(posizione Posizione) string {
	/*
	   La funzione  StringPosizione(posizione Posizione) string  che riceve in input un'instanza di tipo  Posizione  nel
	   parametro  posizione  e restituisce un valore  string  nel formato  etichetta (latitudine, longitudine) , dove
	   etichetta  è il valore  string  che specifica l'etichetta di  posizione , mentre  latitudine  e  longitudine  sono i
	   valori  float64  che specificano rispettivamente la latitudine e la longitudine di  posizione .
	*/
	str := string(posizione.etichetta)
	str += "("
	str += fmt.Sprint(posizione.latitudine)
	str += ","
	str += fmt.Sprint(posizione.longitudine)
	str += ")"
	//+posizione.longitudine
	return str

}

func StringPercorso(percorso []Posizione) string {

	/*
	   La funzione  StringPercorso(percorso []Posizione) string  che riceve in input un'instanza di tipo  []Posizione
	   nel parametro  percorso  e restituisce un valore  string  nel formato  Percorso da POSIZIONE_1 a POSIZIONE_2,
	   cambi: N , dove  POSIZIONE_1  e  POSIZIONE_2  sono le rappresentazioni  string  delle istanze di tipo  Posizione  che
	   rappresentano gli estremi del  percorso , mentre  N  è il numero delle posizioni intermedie (quindi esclusi gli estremi)
	   del  percorso
	*/
	return "Percorso da " + StringPosizione(percorso[0]) + "a " + StringPosizione(percorso[len(percorso)-1]) + ", cambi:" + fmt.Sprint(len(percorso)-1)
}