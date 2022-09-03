/*
Uno dei compiti più importanti di un compilatore è quello di controllare se le parentesi sono ben bilanciate.

Ad ogni parentesi aperta deve corrispondere una parentesi chiusa, e le coppie di parentesi devono essere innestate propriamente.

Un esempio di parentesi tonde ben bilanciate è il seguente: (())()

Un esempio di parentesi tonde non ben bilanciate è il seguente: ())(

Notare che in quest’ultimo esempio, anche se ad ogni parentesi aperta corrisponde una parentesi chiusa, le coppie di parentesi non sono propriamente innestate (viene chiusa una parentesi di troppo ed una parentesi rimane aperta senza mai essere chiusa).

*/

package main

import "fmt"

func main() {
	var serie string

	fmt.Scan(&serie)

	isBalanced(serie)
	//fmt.Println(getSubSring(serie))
	counter := 0

	for i := 0; i < len(serie); i++ {
		for j := i + 1; j <= len(serie); j++ {
			subsequence := serie[i:j]
			if isBalanced(subsequence) {
				counter++
				fmt.Print(counter, ") ", subsequence, "\n")
			}
		}
	}

}

func isBalanced(sequence string) bool {
	countOpen := 0
	for i := 0; i < len(sequence); i++ {
		switch {
		case sequence[i] == 40:
			fmt.Println("(")
			countOpen++
		case sequence[i] == 41:
			fmt.Println(")")
			countOpen--

		default:
			fmt.Println("o")

			return false

		}
		fmt.Println(countOpen)
		if countOpen < 0 {
			fmt.Println("Non bilanciata")
			return false
		}

	}
	if countOpen == 0 {
		fmt.Println("bilanciata")

		return true
	} else {
		fmt.Println("Non bilanciata")

		return false
	}
}

func getSubSring(serie string) []string {
	countOpen := 0
	var subString []string

	for j := 0; j < len(serie); j++ {
		pointerStart := 0
		pointerEnd := 0
		fmt.Println("j:", j, "n", serie[j])
		if serie[j] == 40 {
			for i := j; i < len(serie)-j; i++ {
				switch {
				case serie[i] == 40:
					fmt.Println("(")
					countOpen++
					if pointerStart == 0 && i == j {
						fmt.Println("i:", i, "j", j)
						fmt.Println("setto Start", i)

						pointerStart = i
					}
				case serie[i] == 41:
					fmt.Println(")")
					countOpen--

				}
				if countOpen == 0 {
					fmt.Println("setto End", i)
					pointerEnd = i

					subString = append(subString, serie[pointerStart:pointerEnd+1])

				}

			}
		}
	}

	return subString
}
