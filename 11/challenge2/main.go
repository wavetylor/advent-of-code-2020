package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	occ       = "#"
	empty     = "L"
	floor     = "."
	tolerance = 5
)

func main() {
	dat, _ := ioutil.ReadFile("../in.txt")

	lines := strings.Split(string(dat), "\n")

	num := 1
	seats := createSim(lines)
	stable := false
	fmt.Println("Initial:")
	printer(seats)
	for !stable {
		fmt.Println("Round " + fmt.Sprint(num) + ":")
		seats, stable = simRound(seats)
		printer(seats)
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
		if numOccupied(seats, i, j) >= tolerance {
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
	x, y := i, j
	// top left
	for x > 0 && y > 0 {
		test := seats[x-1][y-1]
		if test != floor {
			if test == occ {
				count++
			}
			break
		}
		x--
		y--
	}
	x, y = i, j

	// top
	for x > 0 {
		test := seats[x-1][y]
		if test != floor {
			if test == occ {
				count++
			}
			break
		}
		x--
	}
	x, y = i, j

	// top right
	for x > 0 && y+1 < cols {
		test := seats[x-1][y+1]
		if test != floor {
			if test == occ {
				count++
			}
			break
		}
		x--
		y++
	}
	x, y = i, j

	// right
	for y+1 < cols {
		test := seats[x][y+1]
		if test != floor {
			if test == occ {
				count++
			}
			break
		}
		y++
	}
	x, y = i, j

	// bottom right
	for x+1 < rows && y+1 < cols {
		test := seats[x+1][y+1]
		if test != floor {
			if test == occ {
				count++
			}
			break
		}
		x++
		y++
	}
	x, y = i, j

	// bottom
	for x+1 < rows {
		test := seats[x+1][y]
		if test != floor {
			if test == occ {
				count++
			}
			break
		}
		x++
	}
	x, y = i, j

	// bottom left
	for x+1 < rows && y-1 >= 0 {
		test := seats[x+1][y-1]
		if test != floor {
			if test == occ {
				count++
			}
			break
		}
		x++
		y--
	}
	x, y = i, j

	// left
	for y-1 >= 0 {
		test := seats[x][y-1]
		if test != floor {
			if test == occ {
				count++
			}
			break
		}
		y--
	}
	return count
}

func shouldSit(seats [][]string, i int, j int) bool {
	should := true
	rows := len(seats)
	cols := len(seats[0])

	x, y := i, j
	// top left
	for x > 0 && y > 0 {
		test := seats[x-1][y-1]
		if test != floor {
			if test == occ {
				should = false
			}
			break
		}
		x--
		y--
	}
	x, y = i, j

	// top
	for should && x > 0 {
		test := seats[x-1][y]
		if test != floor {
			if test == occ {
				should = false
			}
			break
		}
		x--
	}
	x, y = i, j

	// top right
	for should && x > 0 && y+1 < cols {
		test := seats[x-1][y+1]
		if test != floor {
			if test == occ {
				should = false
			}
			break
		}
		x--
		y++
	}
	x, y = i, j

	// right
	for should && y+1 < cols {
		test := seats[x][y+1]
		if test != floor {
			if test == occ {
				should = false
			}
			break
		}
		y++
	}
	x, y = i, j

	// bottom right
	for should && x+1 < rows && y+1 < cols {
		test := seats[x+1][y+1]
		if test != floor {
			if test == occ {
				should = false
			}
			break
		}
		x++
		y++
	}
	x, y = i, j

	// bottom
	for should && x+1 < rows {
		test := seats[x+1][y]
		if test != floor {
			if test == occ {
				should = false
			}
			break
		}
		x++
	}
	x, y = i, j

	// bottom left
	for should && x+1 < rows && y-1 >= 0 {
		test := seats[x+1][y-1]
		if test != floor {
			if test == occ {
				should = false
			}
			break
		}
		x++
		y--
	}
	x, y = i, j

	// left
	for should && y-1 >= 0 {
		test := seats[x][y-1]
		if test != floor {
			if test == occ {
				should = false
			}
			break
		}
		y--
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

func printer(seats [][]string) {

	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			fmt.Print(seats[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Println()
}
