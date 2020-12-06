package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	dat, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(dat), "\n")

	total := 0

	m := make(map[string]bool)
	for _, line := range lines {
		if line == "" {
			total = total + len(m)
			fmt.Println("Total for group: " + fmt.Sprint(len(m)))
			m = make(map[string]bool)
			continue
		}

		for i := 0; i < len(line); i++ {
			m[string(line[i])] = true
		}
	}

	fmt.Println("Last line: " + fmt.Sprint(len(m)))
	total = total + len(m)
	fmt.Println(total)
}
