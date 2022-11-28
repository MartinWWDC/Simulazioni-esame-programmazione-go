package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var nums []int
	temp := os.Args[1:]
	for n := range temp {
		//fmt.Println(temp, "n:", n)
		h, er := strconv.Atoi(temp[n])
		if er == nil {
			if n == 0 {
				nums = append(nums, h)
			} else {
				//fmt.Println("N:", n)
				if n%2 == 0 && h > nums[len(nums)-1] || n%2 != 0 && h < nums[len(nums)-1] {
					nums = append(nums, h)

				} else {
					fmt.Println("Valore in posizione", n, " non valido")
					break
				}

			}
		} else {
			fmt.Println("Valore in posizione", n, " non valido: ", er)
			break
		}

	}
	fmt.Println(nums)

}
