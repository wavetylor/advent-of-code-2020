package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

var col0 = 0

func main() {
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
	fmt.Println("Column 0 count: " + fmt.Sprint(col0))
}

func getID(seat string) int {
	row := bst(seat[:7], "F", 127)
	col := bst(seat[7:], "L", 7)

	fmt.Println("Row: " + fmt.Sprint(row) + ", col: " + fmt.Sprint(col))
	if col == 0 {
		col0++
	}

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
