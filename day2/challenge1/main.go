package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(bytes), "\n")

	validPasswords := 0
	r := regexp.MustCompile(`(?P<first>\d+)\-(?P<second>\d+) (?P<character>\w): (?P<pass>\w+)`)
	names := r.SubexpNames()

	for _, val := range lines {

		result := r.FindAllStringSubmatch(val, -1)
		m := map[string]string{}
		for i, n := range result[0] {
			m[names[i]] = n
		}

		ammount := strings.Count(m["pass"], m["character"])
		lower, _ := strconv.Atoi(m["first"])
		upper, _ := strconv.Atoi(m["second"])
		if lower <= ammount && ammount <= upper {
			validPasswords++
		}
	}

	fmt.Println(validPasswords)

}
