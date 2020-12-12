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

		pass := m["pass"]
		matchChar := m["character"]
		lower, _ := strconv.Atoi(m["first"])
		upper, _ := strconv.Atoi(m["second"])
		char1 := string(pass[lower-1])
		char2 := string(pass[upper-1])

		if (char1 == matchChar || char2 == matchChar) && !(char1 == matchChar && char2 == matchChar) {
			validPasswords++
		}
	}

	fmt.Println(validPasswords)

}
