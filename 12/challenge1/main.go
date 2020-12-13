package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ship struct {
	direction string
	x         int
	y         int
}

var (
	n = "N"
	s = "S"
	e = "E"
	w = "W"
	l = "L"
	r = "R"
	f = "F"
)

var leftMap = map[string]map[int]string{
	n: map[int]string{
		90:  w,
		180: s,
		270: e,
	},
	s: map[int]string{
		90:  e,
		180: n,
		270: w,
	},
	e: map[int]string{
		90:  n,
		180: w,
		270: s,
	},
	w: map[int]string{
		90:  s,
		180: e,
		270: n,
	},
}

var rightMap = map[string]map[int]string{
	s: map[int]string{
		90:  w,
		180: n,
		270: e,
	},
	n: map[int]string{
		90:  e,
		180: s,
		270: w,
	},
	w: map[int]string{
		90:  n,
		180: e,
		270: s,
	},
	e: map[int]string{
		90:  s,
		180: w,
		270: n,
	},
}

func main() {
	dat, _ := ioutil.ReadFile("../in.txt")

	lines := strings.Split(string(dat), "\n")

	s := ship{e, 0, 0}

	for _, line := range lines {
		action := string(line[0])
		move, _ := strconv.Atoi(line[1:])

		if action == l {
			s.turnLeft(move)
		} else if action == r {
			s.turnRight(move)
		} else if action == f {
			s.updatePosition(move)
		} else {
			s.updatePos(action, move)
		}

		fmt.Println(s.direction)
		fmt.Printf("Coords: (%d,%d)\n", s.x, s.y)
	}
	fmt.Printf("%d\n", s.dist())
}

func (o *ship) turnLeft(degree int) {
	newDir, _ := leftMap[o.direction][degree]
	o.direction = newDir
}

func (o *ship) turnRight(degree int) {
	newDir, _ := rightMap[o.direction][degree]
	o.direction = newDir
}

func (o *ship) updatePosition(moves int) {
	o.updatePos(o.direction, moves)
}

func (o *ship) updatePos(dir string, moves int) {
	switch dir {
	case n:
		o.y += moves
	case s:
		o.y -= moves
	case e:
		o.x += moves
	case w:
		o.x -= moves
	}
}

func (o *ship) dist() int {
	x := o.x
	y := o.y
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}
