package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ship struct {
	x int
	y int
}

type waypoint struct {
	x int
	y int
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

func main() {
	dat, _ := ioutil.ReadFile("../in.txt")

	lines := strings.Split(string(dat), "\n")

	s := ship{0, 0}
	w := waypoint{10, 1}

	for _, line := range lines {
		action := string(line[0])
		moves, _ := strconv.Atoi(line[1:])

		if action == f {
			s.updatePosition(moves, w)
		} else if action == l {
			w.rotateLeft(moves)
		} else if action == r {
			w.rotateRight(moves)
		} else {
			w.update(action, moves)
		}

		fmt.Printf("Ship Coords:     (%d,%d)\n", s.x, s.y)
		fmt.Printf("WayPoint Coords: (%d,%d)\n", w.x, w.y)
	}
	fmt.Printf("%d\n", s.dist())
}

func (z *waypoint) rotateLeft(deg int) {
	times := deg / 90
	for times > 0 {
		x := z.x
		y := z.y

		z.x = y * -1
		z.y = x
		times--
	}
}

func (z *waypoint) rotateRight(deg int) {
	times := deg / 90
	for times > 0 {
		x := z.x
		y := z.y

		z.x = y
		z.y = x * -1
		times--
	}
}

func (z *waypoint) update(dir string, moves int) {
	switch dir {
	case n:
		z.y += moves
	case s:
		z.y -= moves
	case e:
		z.x += moves
	case w:
		z.x -= moves
	}
}

func (o *ship) updatePosition(moves int, w waypoint) {
	o.x = o.x + w.x*moves
	o.y = o.y + w.y*moves
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
