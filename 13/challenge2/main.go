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
	fmt.Println("Solving for: " + lines[1])

	ids, count := buses(lines[1])

	syncBusIdx := 1
	syncBus := 1
	currentLcm := ids[0]
	t := currentLcm

	for true {
		temp := t + syncBusIdx
		for i, id := range ids[syncBusIdx:] {
			if id > 0 {
				if temp%id == 0 {
					currentLcm *= id
					syncBus++
					temp++
					syncBusIdx = syncBusIdx + i + 1
				} else {
					break
				}
			} else {
				temp++
			}
		}
		if syncBus == count {
			break
		}
		t += currentLcm
	}
	fmt.Printf("Earliest timestamp: %d\n", t)
}

func buses(in string) ([]int, int) {
	ret := make([]int, 0)
	ids := strings.Split(in, ",")
	i := 0
	for _, b := range ids {
		if b != "x" {
			val, _ := strconv.Atoi(string(b))
			ret = append(ret, val)
			i++
		} else {
			ret = append(ret, 0)
		}
	}
	return ret, i
}
