package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("../in.txt")

	lines := strings.Split(string(dat), "\n")

	time, _ := strconv.Atoi(lines[0])
	ids := buses(lines[1])

	smallestDiff := 10000
	winnerID := -1
	for _, b := range ids {
		wait := b - (time % b)
		if wait < smallestDiff {
			smallestDiff = wait
			winnerID = b
		}
	}
	fmt.Printf("Smallest wait %d, bus id: %d. Value: %d\n", smallestDiff, winnerID, smallestDiff*winnerID)
}

func buses(in string) []int {
	ret := make([]int, 0)
	ids := strings.Split(in, ",")
	for _, b := range ids {
		if b != "x" {
			val, _ := strconv.Atoi(string(b))
			ret = append(ret, val)
		}
	}
	return ret
}
