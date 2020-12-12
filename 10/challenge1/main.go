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

	jolts := make([]int, len(lines))
	for i, l := range lines {
		x, _ := strconv.Atoi(l)
		jolts[i] = x
	}
	sort.Ints(jolts)

	voltsOne := 0
	voltsTwo := 0
	voltsThree := 0

	prev := 0
	start := 0
	for i := 1; i < len(jolts); i++ {
		if prev == 0 {
			diff := jolts[0] - start
			switch diff {
			case 1:
				voltsOne++
			case 2:
				voltsTwo++
			case 3:
				voltsThree++
			default:
				fmt.Println("Oops!")
			}
			start += jolts[0]
			prev = i
			continue
		}

		diff := jolts[i] - jolts[prev]
		switch diff {
		case 1:
			voltsOne++
		case 2:
			voltsTwo++
		case 3:
			voltsThree++
		default:
			fmt.Println("Oops!")
		}

		prev = i
	}

	fmt.Printf("Volt Diff:\n\t(1): %d\n\t(2): %d\n\t(3): %d\n", voltsOne, voltsTwo, voltsThree)
	fmt.Println("Answer: " + fmt.Sprintf("%d", (voltsOne+1)*(voltsThree+1)))
}
