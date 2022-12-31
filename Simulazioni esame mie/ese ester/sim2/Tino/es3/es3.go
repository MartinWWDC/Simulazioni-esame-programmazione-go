package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type data struct {
	data string
	typ  string
	mov  int
}

func main() {
	dat, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(dat)
	//fmt.Println(string(dat))
	arrD := []data{}
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		date := strings.Split(arr[0], "_")
		if len(date[1]) == 1 {
			date[1] = "0" + date[1]
		}
		if len(date[2]) == 1 {
			date[2] = "0" + date[2]
		}
		i, _ := strconv.Atoi(arr[2])
		arrD = append(arrD, data{date[0] + "-" + date[1] + "-" + date[2], arr[1], i})
	}
	fmt.Println(arrD)
	mapp := make(map[string]int)
	for _, i := range arrD {
		if i.typ == "P" {
			mapp[i.data] -= i.mov
		} else {
			mapp[i.data] += i.mov

		}
	}
	fmt.Println(mapp)
}
