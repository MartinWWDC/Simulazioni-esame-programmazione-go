package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(LeggiTesto())
}

func LeggiTesto() string {
	text := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}
	return text
}
