package main

import (
	"fmt"
	"os"
)

func main() {
	sequenza := os.Args[1:]
	mappa := GeneraSottosequenze(sequenza)
	delete(mappa, "")
	fmt.Println(mappa)

}

func GeneraSottosequenze(sequenza []string) map[string]int {
	fmt.Println(sequenza)

	sottosequent := make(map[string]int)
	var seq, primo, ultimo string
	sequenza = findAndDelete(sequenza, " ")
	for i := 0; i < len(sequenza)+1; i++ {
		for j := i; j < len(sequenza)+1; j++ {
			seq = ""
			subSeq := sequenza[i:j]
			//fmt.Println("i:", i, "j:", j)
			//fmt.Println("sequenza:", subSeq)

			for item := range subSeq {
				//fmt.Println("item id:", item, "item", subSeq[item])
				seq += (subSeq[item])
			}
			//fmt.Println(seq)
			//fmt.Println(seq[len(seq)])
			//fmt.Println("item:", string(seq[0]))

			for i := 0; i < len(seq); i++ {
				//fmt.Println("item:", string(seq[0]))
				//fmt.Println("item finale:", string(seq[len(seq)-1]))
				primo = string(seq[0])
				ultimo = string(seq[len(seq)-1])

			}
			if primo == ultimo {
				if _, ok := sottosequent[seq]; ok {
					//fmt.Println("c'è")
					sottosequent[seq] += 1

				} else {
					//fmt.Println("non c'è")
					sottosequent[seq] = 1
				}

			}

		}
	}
	for item := range sottosequent {
		if len(item) < 3 {
			delete(sottosequent, item)
		}
	}
	return sottosequent
}

func findAndDelete(s []string, item string) []string {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}
