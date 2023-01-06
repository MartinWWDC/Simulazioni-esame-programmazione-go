package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	eq1 := os.Args[1]
	eq2 := os.Args[2]
	m1 := getM(eq1)
	m2 := getM(eq2)
	if m2-m1 <= 0.0001 {
		fmt.Println("parall")
	}

}
func getM(eq string) float64 {
	arr := strings.Split(eq, "x")
	a, _ := strconv.Atoi(arr[0])
	b, _ := strconv.Atoi(strings.Split(arr[1], "y")[0])
	return -1 * float64(a/b)
}
