package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("in.txt")
	lines := strings.Split(string(dat), "\n")
	size := 25
	nums := readPreamble(lines, size)

	for i := size; i < len(lines); i++ {
		x, _ := strconv.Atoi(lines[i])
		valid := false

		for j := 0; j < len(nums); j++ {
			for p := j + 1; p < len(nums); p++ {
				if nums[j]+nums[p] == x {
					valid = true
				}
			}
		}
		if !valid {
			fmt.Printf("Not valid: %d\n", x)
		}
		nums[i%size] = x
	}
}

func readPreamble(lines []string, preambleSize int) []int {
	nums := make([]int, preambleSize)

	for i := 0; i < preambleSize; i++ {
		x, _ := strconv.Atoi(lines[i])
		nums[i] = x
	}

	return nums
}
