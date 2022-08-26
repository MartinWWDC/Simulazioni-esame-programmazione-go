package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//eq := os.Args[1]
	eq := "-5000xˆ2-1x+2=0"
	estrai_a := strings.Split(eq, "x^2") // divide la stringa in a e quello che viene dopo "x^2"
	a, _ := strconv.ParseFloat(estrai_a[0], 64)

	estrai_b := strings.Split(estrai_a[1], "x") // divide la stringa RIMANENTE in b e quello che viene dopo "x"
	b, _ := strconv.ParseFloat(estrai_b[0], 64)

	estrai_c := strings.Split(estrai_b[1], "=") // divide la stringa RIMANENTE in c e quello che viene dopo "="
	c, _ := strconv.ParseFloat(estrai_c[0], 64)

	discriminante := b*b - (4 * a * c)

	fmt.Println(c)
	fmt.Println(discriminante)

	/*
		for i := 0; i < len(sep); i++ {

			fmt.Println(sep[i])
			item := sep[i]
			for j := 0; j < len(sep[i]); j++ {
				fmt.Println("char", string(item[j]))
				if string(item[j]) == "x" || string(item[j]) == "=" {
					priv := item[0:j]
					fmt.Println(priv)
					items = append(items, priv)
					j = len(sep[i])
				}
			}

		}*/
	//fmt.Println(items)
}
