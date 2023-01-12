package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(normalize("aaaa_m_g"))

	scanner := bufio.NewScanner(os.Stdin)
	mapp := make(map[string]int)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		date := normalize(arr[0])
		k, _ := strconv.Atoi(arr[2])
		if arr[1] == "V" {
			mapp[date] += k
		} else {
			mapp[date] -= k

		}

	}
	fmt.Println(mapp)
}
func normalize(date string) string {
	arr := strings.Split(date, "_")
	if len(arr[1]) < 2 {
		arr[1] = "0" + arr[1]
	}
	if len(arr[2]) < 2 {
		arr[2] = "0" + arr[2]
	}
	return arr[0] + "/" + arr[1] + "/" + arr[2]
}
