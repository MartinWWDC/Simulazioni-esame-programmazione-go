package main

import (
	"fmt"
	"os"
)

func main() {
	parentesi := os.Args[1]
	getSub(parentesi)

}

func isBalanced(sequence string) bool {
	counter := 0
	for _, i := range sequence {
		if i == '(' {
			counter++
		}
		if i == ')' {
			counter--
		}
		if counter < 0 {
			return false
		}
	}
	if counter == 0 {
		return true
	} else {
		return false
	}
}
func getSub(sequence string) {
	if isBalanced(sequence) {
		fmt.Println("bilanciata")
		fmt.Println("----")
		fmt.Println("Sottosequenze bilanciate:")
		for i := range sequence {
			for h := i; h <= len(sequence); h++ {
				if len(sequence[i:h]) > 1 && isBalanced(sequence[i:h]) {
					fmt.Println(sequence[i:h])
				}
			}
		}
	} else {
		fmt.Println("non bilanciata")

	}

}
