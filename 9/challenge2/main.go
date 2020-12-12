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
	invalid := findInvalid(lines, nums)
	magic := computeMagicNumber(lines, invalid)
	fmt.Println(magic)

}

func computeMagicNumber(lines []string, toSum int) int {
	for i := 0; i < len(lines); i++ {
		runningTotal := conv(lines[i])
		min := runningTotal
		max := runningTotal
		for j := i + 1; j < len(lines); j++ {
			test := conv(lines[j])
			if runningTotal+test < toSum {
				runningTotal += test
				if test > max {
					max = test
				} else if test < min {
					min = test
				}
			} else if runningTotal+test == toSum {
				fmt.Printf("Found combo.\n")
				if test > max {
					max = test
				} else if test < min {
					min = test
				}
				return max + min
			}
		}
	}
	return -1
}

func findInvalid(lines []string, nums []int) int {
	size := len(nums)
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
			return x
		}
		nums[i%size] = x
	}
	return -1
}

func readPreamble(lines []string, preambleSize int) []int {
	nums := make([]int, preambleSize)

	for i := 0; i < preambleSize; i++ {
		x, _ := strconv.Atoi(lines[i])
		nums[i] = x
	}

	return nums
}

func conv(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
