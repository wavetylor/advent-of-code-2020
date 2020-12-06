package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var byr = regexp.MustCompile(`byr:(\d{4})`)
var iyr = regexp.MustCompile(`iyr:(\d{4})`)
var eyr = regexp.MustCompile(`eyr:(\d{4})`)
var hgt = regexp.MustCompile(`hgt:(\d{2,3})(cm|in)`)
var hcl = regexp.MustCompile(`hcl:#([a-f0-9]{6})`)
var ecl = regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)`)
var pid = regexp.MustCompile(`pid:(\d{9})($|\s)`)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(dat), "\n")

	valid := 0

	toGo := 7
	for _, line := range lines {
		if line == "" || toGo == 0 {
			fmt.Println("resetting")
			toGo = 7
			continue
		}

		toGo = has(line, toGo, byrTest)
		toGo = has(line, toGo, iyrTest)
		toGo = has(line, toGo, eyrTest)
		toGo = has(line, toGo, hgtTest)
		toGo = has(line, toGo, hclTest)
		toGo = has(line, toGo, eclTest)
		toGo = has(line, toGo, pidTest)

		if toGo == 0 {
			fmt.Println("Matched: " + line)
			valid++
		}
	}

	fmt.Println("Valid: " + fmt.Sprint(valid))
}

func has(line string, counter int, tester func(l string) bool) int {
	if tester(line) {
		fmt.Println("Passed: " + line)
		return counter - 1
	}
	return counter
}

func byrTest(l string) bool {
	matches := byr.FindStringSubmatch(l)
	if len(matches) == 2 {
		year, _ := strconv.Atoi(matches[1])
		return 1920 <= year && year <= 2002
	}
	return false
}

func iyrTest(l string) bool {
	matches := iyr.FindStringSubmatch(l)
	if len(matches) == 2 {
		year, _ := strconv.Atoi(matches[1])
		return 2010 <= year && year <= 2020
	}
	return false
}

func eyrTest(l string) bool {
	matches := eyr.FindStringSubmatch(l)
	if len(matches) == 2 {
		year, _ := strconv.Atoi(matches[1])
		return 2020 <= year && year <= 2030
	}
	return false
}

func hgtTest(l string) bool {
	matches := hgt.FindStringSubmatch(l)
	if len(matches) == 3 {
		num, _ := strconv.Atoi(matches[1])
		inCm := matches[2]

		if inCm == "cm" {
			return 150 <= num && num <= 193
		}
		return 59 <= num && num <= 76

	}
	return false
}

func hclTest(l string) bool {
	return len(hcl.FindStringSubmatch(l)) == 2
}

func eclTest(l string) bool {
	return len(ecl.FindStringSubmatch(l)) == 2
}

func pidTest(l string) bool {
	return len(pid.FindStringSubmatch(l)) == 3
}
