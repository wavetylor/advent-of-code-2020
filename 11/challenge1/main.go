package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	occ   = "#"
	empty = "L"
	floor = "."
)

func main() {
	dat, _ := ioutil.ReadFile("../in.txt")

	lines := strings.Split(string(dat), "\n")

	num := 0
	seats := createSim(lines)
	stable := false
	for !stable {
		seats, stable = simRound(seats)
		num++
	}
	total := countOcc(seats)

	fmt.Println(total)
}

func createSim(lines []string) [][]string {
	sim := make([][]string, len(lines))
	for i := range sim {
		sim[i] = make([]string, len(lines[0]))
	}

	for i, l := range lines {
		for j, c := range strings.Split(l, "") {
			sim[i][j] = c
		}
	}
	return sim
}

func simRound(seats [][]string) ([][]string, bool) {
	next := newSeats(len(seats), len(seats[0]))

	for i, row := range seats {
		for j := range row {
			next[i][j] = getChange(seats, i, j)
		}
	}

	return next, isStable(next, seats)
}

func getChange(seats [][]string, i int, j int) string {
	seat := seats[i][j]
	if seat == floor {
		return floor
	} else if seat == occ {
		if numOccupied(seats, i, j) >= 4 {
			return empty
		}
		return occ
	} else {
		if shouldSit(seats, i, j) {
			return occ
		}
		return empty
	}
}

func newSeats(row int, col int) [][]string {
	sim := make([][]string, row)
	for i := range sim {
		sim[i] = make([]string, col)
	}
	return sim
}

func numOccupied(seats [][]string, i int, j int) int {
	count := 0
	rows := len(seats)
	cols := len(seats[0])
	// top left
	if i > 0 && j > 0 {
		if seats[i-1][j-1] == occ {
			count++
		}
	}

	// top
	if i > 0 {
		if seats[i-1][j] == occ {
			count++
		}
	}

	// top right
	if i > 0 && j+1 < cols {
		if seats[i-1][j+1] == occ {
			count++
		}
	}

	// right
	if j+1 < cols {
		if seats[i][j+1] == occ {
			count++
		}
	}

	// bottom right
	if i+1 < rows && j+1 < cols {
		if seats[i+1][j+1] == occ {
			count++
		}
	}

	// bottom
	if i+1 < rows {
		if seats[i+1][j] == occ {
			count++
		}
	}

	// bottom left
	if i+1 < rows && j-1 >= 0 {
		if seats[i+1][j-1] == occ {
			count++
		}
	}

	// left
	if j-1 >= 0 {
		if seats[i][j-1] == occ {
			count++
		}
	}

	return count

}

func shouldSit(seats [][]string, i int, j int) bool {
	should := true
	rows := len(seats)
	cols := len(seats[0])

	if i > 0 && j > 0 {
		if seats[i-1][j-1] == occ {
			should = false
		}
	}

	// top
	if should && i > 0 {
		if seats[i-1][j] == occ {
			should = false
		}
	}

	// top right
	if should && i > 0 && j+1 < cols {
		if seats[i-1][j+1] == occ {
			should = false
		}
	}

	// right
	if should && j+1 < cols {
		if seats[i][j+1] == occ {
			should = false
		}
	}

	// bottom right
	if should && i+1 < rows && j+1 < cols {
		if seats[i+1][j+1] == occ {
			should = false
		}
	}

	// bottom
	if should && i+1 < rows {
		if seats[i+1][j] == occ {
			should = false
		}
	}

	// bottom left
	if should && i+1 < rows && j-1 >= 0 {
		if seats[i+1][j-1] == occ {
			should = false
		}
	}

	// left
	if should && j-1 >= 0 {
		if seats[i][j-1] == occ {
			should = false
		}
	}
	return should
}

func isStable(new [][]string, old [][]string) bool {
	for i := 0; i < len(new); i++ {
		for j := 0; j < len(new[i]); j++ {
			if new[i][j] != old[i][j] {
				return false
			}
		}
	}
	return true
}

func countOcc(seats [][]string) int {
	count := 0
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == occ {
				count++
			}
		}
	}
	return count
}
