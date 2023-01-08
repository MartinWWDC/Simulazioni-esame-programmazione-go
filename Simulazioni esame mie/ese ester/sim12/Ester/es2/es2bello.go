package main

import (
	. "fmt"
	"sort"
)

func subb(s string) map[string]int {
	mappa := make(map[string]int)
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			sott := s[i:j]
			if len(sott) >= 3 && sott[0] == sott[len(sott)-1] {
				mappa[sott]++
			}
		}
	}
	return mappa
}

func main() {
	s := "abbabba"
	m := subb(s)
	var sta []string

	for i := range m {
		sta = append(sta, i)
	}

	sort.SliceStable(sta, func(i, j int) bool { return len(sta[i]) > len(sta[j]) })

	for i := 0; i < len(sta); i++ {
		Println(sta[i], m[sta[i]])
	}

}
