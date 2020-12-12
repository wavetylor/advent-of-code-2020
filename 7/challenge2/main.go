package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type graph struct {
	nodes map[string]struct{}
	edges map[string][]bag
}

type bag struct {
	color string
	count int
}

func (g *graph) addNode(n string) {
	g.nodes[n] = struct{}{}
}

func (g *graph) addEdge(src, dest string, count int) {
	if g.edges == nil {
		g.edges = make(map[string][]bag)
	}
	g.edges[src] = append(g.edges[src], bag{dest, count})
}

func main() {

	dat, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(dat), "\n")

	g := buildGraph(lines)
	fmt.Println("Shiny Bag in: " + fmt.Sprint(countShiny(g)))

}

func countShiny(g graph) int {
	return hasGold("shiny gold", g) - 1
}

func hasGold(n string, g graph) int {
	runningSum := 1
	for _, e := range g.edges[n] {
		runningSum = runningSum + e.count*hasGold(e.color, g)
	}
	return runningSum
}

func buildGraph(lines []string) graph {
	var g graph
	g.nodes = make(map[string]struct{})

	var edgeMatcher = regexp.MustCompile(`(\d+) (.+) bag(s?)`)
	for _, line := range lines {
		parts := strings.Split(line, " bags contain ")

		n := parts[0]
		g.addNode(n)

		edges := strings.Split(parts[1], ", ")

		for _, edge := range edges {
			matches := edgeMatcher.FindStringSubmatch(edge)
			if len(matches) > 0 {
				count, _ := strconv.Atoi(matches[1])
				color := matches[2]
				if _, ok := g.nodes[color]; !ok {
					g.addNode(color)
				}
				g.addEdge(n, color, count)
			}
		}
	}
	return g
}
