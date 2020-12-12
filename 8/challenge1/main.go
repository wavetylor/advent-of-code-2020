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
	acc := 0
	visited := make(map[int]struct{})

	for i := 0; i < len(lines); i++ {
		if _, ok := visited[i]; ok {
			fmt.Println("visited: " + fmt.Sprint(i) + " twice!")
			break
		}
		visited[i] = struct{}{}
		line := lines[i]

		if line[:3] == "acc" {
			x, _ := strconv.Atoi(line[4:])
			acc += x
		} else if line[:3] == "jmp" {
			x, _ := strconv.Atoi(line[4:])
			i = x + i - 1
		}
	}

	fmt.Println("Acc final Value: " + fmt.Sprint(acc))
}
