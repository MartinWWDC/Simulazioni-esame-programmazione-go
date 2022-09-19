package main

func main() {

}

/*
i= numerone pieno di valori int
j= numero del quale voglio sapere quanto volte  è all'interno di i
*/

func check(i int, j int) int {
	if i == 0 {
		return 0
	}
	if i%10 == j {
		return check(i/10, j) + 1
	}
	return check(i/10, j)
}
