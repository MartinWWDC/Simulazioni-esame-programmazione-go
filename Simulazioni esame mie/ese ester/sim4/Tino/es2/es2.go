package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Comando struct {
	direzione string
	passi     int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mapp := make(map[string]int)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(arr[1])
		mapp[arr[0]] += n
	}
	fmt.Println(mapp)
}
