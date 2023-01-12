package main

func main() {

}

func isGood(str string) (bool, string) {
	c := true
	err := ""
	if !(len(str) == 8) {
		c = false
		err += "- La pw deve avere una lunghezza di 8 caratteri \n"
	}
	nMin := 0
	nMaius := 0
	nNums := 0

	for _, h := range str {
		j := false
		if h >= 'A' && h <= 'Z' {
			nMaius++
			j = true
		}
		if h >= 'a' && h <= 'z' {
			nMin++
			j = true

		}
		if h >= '1' && h <= '9' {
			nNums++
			j = true

		}
		if j {
			c = false
			err += "- Tutti i caratteri della pw devono rappresentare lettere o cifre decimali\n"
		}
	}
	if nMin < 2 {
		c = false
		err += "- Almeno 2 caratteri della pw devono rappresentare delle lettere minuscole\n"
	}
	if nMaius > 3 {
		c = false
		err += "- Al massino 3 caratteri della pw devono rappresentare delle lettere minuscole\n"
	}
	if nNums < 2 {
		c = false
		err += "- Almeno 2 caratteri della pw devono rappresentare delle cifre decimali\n"
	}
	if !c {
		err = "La pw non è definita correttamente:\n		" + err
	}
	return c, err
}
