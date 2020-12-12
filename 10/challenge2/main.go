package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("../in.txt")
	lines := strings.Split(string(dat), "\n")
	jolts := sorted(lines)
	max := jolts[len(jolts)-1]

	counter := make([]int, max+1)
	counter[0] = 1
	for _, j := range jolts {
		total := 0
		for i := 1; i < 4; i++ {
			if j-i >= 0 && j-i < len(counter)-1 {
				total += counter[j-i]
			}
		}
		counter[j] = total
	}
	fmt.Println(counter[jolts[len(jolts)-1]])
}

func sorted(lines []string) []int {
	jolts := make([]int, len(lines))
	for i, l := range lines {
		x, _ := strconv.Atoi(l)
		jolts[i] = x
	}
	sort.Ints(jolts)
	return jolts
}
