package main

import "fmt"

func main() {
	stringa := ""
	var subs string
	var str string
	mapp := make(map[string]int)
	fmt.Scan(&stringa)
	//fmt.Println(stringa)

	for i := 0; i < len(stringa); i++ {
		for j := i; j < len(stringa)+1; j++ {
			subs = stringa[i:j]
			//fmt.Println(subs)
			if len(subs) > 1 {
				//fmt.Println("dentro")
				for y := 0; y < len(subs); y++ {
					//fmt.Println("dentro", subs)
					str += string(subs[len(subs)-y-1])

				}
				if subs == str {

					if _, ok := mapp[str]; ok {
						mapp[str]++
					} else {
						mapp[str] = 1
					}
				}
			}
			str = ""

		}
	}
	fmt.Println(mapp)
}
