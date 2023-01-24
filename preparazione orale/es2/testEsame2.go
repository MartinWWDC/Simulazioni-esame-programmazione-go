package main

/**
 * prendo una slice di interi e resituisce una slice di interi senza ripetizioni e stampare l'occorrenza
 */

func RimuoviRipetizioni(n []int) map[int]int {
	mapp := make(map[int]int)

	for _, i := range n {
		mapp[i]++
	}
	return mapp
}
