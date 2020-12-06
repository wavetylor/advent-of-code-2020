package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\n")

	total := run(lines, 1, 1) * run(lines, 1, 3) * run(lines, 1, 5) * run(lines, 1, 7) * run(lines, 2, 1)
	fmt.Println(total)

}

func run(lines []string, down int, right int) int {
	mapLen := len(lines[0])
	toTheRight := 0
	trees := 0
	for i, line := range lines {

		if i > 0 && i%down == 0 {
			toTheRight += right
			if string(line[toTheRight%mapLen]) == "#" {
				trees++
			}
		}
	}
	fmt.Printf("Hit %d trees, with right: %d, down: %d\n", trees, right, down)
	return trees
}
