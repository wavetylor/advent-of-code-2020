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
	found := false
	acc := 0
	for i, l := range lines {
		if strings.Contains(l, "jmp") || strings.Contains(l, "nop") {
			found, acc = test(i, lines)
			if found {
				break
			}
		}
	}

	fmt.Println("Acc final Value: " + fmt.Sprint(acc))
}

func test(toChange int, lines []string) (bool, int) {
	fmt.Printf("[Test: %d] Changing instruction %s\n", toChange, lines[toChange])
	acc := 0
	finished := true
	visited := make(map[int]struct{})

	for i := 0; i < len(lines); i++ {
		if _, ok := visited[i]; ok {
			fmt.Printf("[Test: %d] Visited: %d twice\n", toChange, i)
			return false, -1
		}
		visited[i] = struct{}{}
		line := lines[i]

		if line[:3] == "acc" {
			x, _ := strconv.Atoi(line[4:])
			acc += x
		} else if line[:3] == "jmp" {
			if i == toChange {
				fmt.Printf("[Test: %d] Ignoring jmp: %s at line: %d\n", toChange, line, i)
				continue
			}
			x, _ := strconv.Atoi(line[4:])
			i = x + i - 1

		} else if line[:3] == "nop" && i == toChange {
			fmt.Printf("[Test: %d] Changing nop to jmp: %s at line: %d\n", toChange, line, i)
			x, _ := strconv.Atoi(line[4:])
			i = x + i - 1
		}
	}
	return finished, acc
}
