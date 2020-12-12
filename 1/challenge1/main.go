package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	dat, err := ioutil.ReadFile("./expenses.txt")
	if err != nil {
		panic(err)
	}

	slices := bytes.Split(dat, []byte("\n"))

	found := false
	for _, num := range slices {
		if found == true {
			break
		}
		for _, num2 := range slices {
			if found == true {
				break
			}
			// fmt.Printf("Found the numbers: %s + %s\n", num, num2)
			int1, _ := strconv.Atoi(fmt.Sprintf("%s", num))

			int2, _ := strconv.Atoi(fmt.Sprintf("%s", num2))

			// fmt.Printf("Found the numbers: %d \n", int1+int2)
			if int1+int2 == 2020 {
				fmt.Printf("Found the numbers: %d, %d\n", int1, int2)
				fmt.Println(int1 * int2)
				found = true
			}
		}
	}

}
