package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	nums, n := leggiNumeri()
	if n && validate(nums) {
		fmt.Println("Sequenza valida.")
	}
}

func leggiNumeri() ([]int, bool) {
	var numeri []int
	check := true
	for i, g := range os.Args {
		if i != 0 {
			if n, error := strconv.Atoi(g); error == nil {
				numeri = append(numeri, (n))
			} else {
				fmt.Println("Valore in posizione ", i, " non valido.")
				//fmt.Println(g)
				check = false
				break
			}
		}
	}
	return numeri, check
}
func validate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if i != 0 {

			if i%2 == 0 {
				if nums[i] > nums[i-1] {

				} else {
					fmt.Println("Valore in posizione ", i, " non valido.")

					return false
				}

			} else {
				if nums[i] < nums[i-1] {

				} else {
					fmt.Println("Valore in posizione ", i, " non valido.")

					return false

				}

			}
		}
	}
	return true
}
