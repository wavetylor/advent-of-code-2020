package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("./expenses.txt")
	if err != nil {
		panic(err)
	}

	slices := strings.Split(string(dat), "\n")

	found := false
	for _, num := range slices {
		if found == true {
			break
		}
		for _, num2 := range slices {
			if found == true {
				break
			}
			int1, _ := strconv.Atoi(num)

			int2, _ := strconv.Atoi(num2)

			if int1+int2 < 2020 {
				for _, num3 := range slices {
					int3, _ := strconv.Atoi(num3)

					if int1+int2+int3 == 2020 {
						fmt.Printf("Found the numbers: %d, %d, %d\n", int1, int2, int3)
						fmt.Println(int1 * int2 * int3)
						found = true
						break
					}

				}
			}

		}
	}

}
