package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])
	d, _ := strconv.Atoi(os.Args[3])
	fmt.Println(N, k, d)
	Arr := GeneraNumeri(N, k)
	fmt.Println(FiltraNumeri(Arr, d))

}
func GeneraNumeri(N, k int) (arr []int) {
	NStr := strconv.Itoa(N)
	for i := range NStr[:len(NStr)-k+1] {
		sub, _ := strconv.Atoi(NStr[:i] + NStr[i+k:])
		arr = append(arr, sub)

	}

	return
}
func FiltraNumeri(sl []int, d int) (arr []int) {
	for _, h := range sl {
		if h%d != 0 {
			arr = append(arr, h)
		}
	}
	return
}
