package main

import "fmt"

func main() {
	n := 4
	L := 3
	H := 3
	u := 0
	c := true
	c2 := true
	for i := 0; i < n+n+1; i++ {
		if i%2 != 0 {
			for j := 0; j < H-2; j++ {
				if c2 {
					for h := 0; h < L; h++ {
						if h == L-1 {
							fmt.Println(u)
							u = add(u)
						} else {
							fmt.Print(" ")
						}
					}

				} else {
					fmt.Println(u)
					u = add(u)
				}

			}
			if c2 {
				c2 = false

			} else {
				c2 = true

			}

		} else {
			arr := []int{}
			for j := 0; j < L; j++ {
				if c {
					fmt.Print(u)
					u = add(u)

				} else {
					arr = append(arr, u)
					u = add(u)
				}

			}
			if c {
				c = false

			} else {
				c = true
				for i := len(arr) - 1; i >= 0; i-- {
					fmt.Print(arr[i])
				}
			}
			fmt.Println("")
		}
	}
}
func add(n int) int {
	if n == 9 {
		return 0
	} else {
		return n + 1
	}
}
