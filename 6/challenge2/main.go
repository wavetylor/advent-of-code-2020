package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var m = make(map[string]int)

func main() {

	dat, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(dat), "\n")

	lineCounter := 0
	total := 0

	for _, line := range lines {
		if line == "" {
			groupTotal := 0
			for _, v := range m {
				if v == lineCounter {
					groupTotal++
				}
			}
			fmt.Println("Total for group: " + fmt.Sprint(groupTotal))

			total = total + groupTotal
			m = make(map[string]int)
			lineCounter = 0
			continue
		}
		lineCounter++
		for i := 0; i < len(line); i++ {
			set(string(line[i]))
		}
	}

	groupTotal := 0
	for _, v := range m {
		if v == lineCounter {
			groupTotal++
		}
	}
	fmt.Println("Last line: " + fmt.Sprint(groupTotal))
	total = total + groupTotal
	fmt.Println(total)
}

func set(key string) {
	if v, ok := m[key]; ok {
		m[key] = v + 1
	} else {
		m[key] = 1
	}
}
