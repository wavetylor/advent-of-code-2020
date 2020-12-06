package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

var seats = [128][8]bool{}

func main() {

	initSeats()

	dat, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(dat), "\n")

	highestID := 0

	for _, seat := range lines {
		id := getID(seat)
		if id > highestID {
			highestID = id
		}
	}

	fmt.Println("Highest id: " + fmt.Sprint(highestID))
	fmt.Println("My seat: " + fmt.Sprint(getSeat()))
}

func getID(seat string) int {
	row := bst(seat[:7], "F", 127)
	col := bst(seat[7:], "L", 7)

	fmt.Println("Row: " + fmt.Sprint(row) + ", col: " + fmt.Sprint(col))
	seats[row][col] = true
	return row*8 + col
}

func bst(toParse string, lowChar string, high int) int {
	lower := 0
	upper := high

	for i := 0; i < len(toParse); i++ {
		m := int(math.Ceil(float64(lower+upper) / float64(2)))

		div := string(toParse[i])
		if div == lowChar {
			upper = m - 1
		} else {
			lower = m
		}
	}
	return lower
}

func initSeats() {
	for x, col := range seats {
		for y := range col {
			seats[x][y] = false
		}
	}
}

func getSeat() int {
	for x, col := range seats {
		for y := range col {
			if x > 0 && !seats[x][y] {
				return x*8 + y
			}
		}
	}
	return -100
}
