package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type graph struct {
	nodes map[string]struct{}
	edges map[string][]string
}

func (g *graph) addNode(n string) {
	g.nodes[n] = struct{}{}
}

func (g *graph) addEdge(src, dest string) {
	if g.edges == nil {
		g.edges = make(map[string][]string)
	}
	g.edges[src] = append(g.edges[src], dest)
}

func main() {

	dat, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(dat), "\n")

	g := buildGraph(lines)
	fmt.Println("Shiny Bag in: " + fmt.Sprint(countShiny(g)))

}

var shiny = 0

func countShiny(g graph) int {

	for n := range g.nodes {
		visited := make(map[string]bool, len(g.nodes))
		for k := range g.nodes {
			visited[k] = false
		}
		if hasGold(n, g, visited) {
			shiny++
		}
	}

	return shiny
}

func hasGold(n string, g graph, visited map[string]bool) bool {
	visited[n] = true

	if contains(g.edges[n], "shiny gold") {
		return true
	}

	for _, e := range g.edges[n] {
		if !visited[e] && e != "shiny gold" {
			if hasGold(e, g, visited) {
				return true
			}
		}
	}
	return false
}

func contains(list []string, term string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == term {
			return true
		}
	}
	return false
}

func buildGraph(lines []string) graph {
	var g graph
	g.nodes = make(map[string]struct{})

	var edgeMatcher = regexp.MustCompile(`\d+ (.+) bag(s?)`)
	for _, line := range lines {
		parts := strings.Split(line, " bags contain ")

		n := parts[0]
		g.addNode(n)

		edges := strings.Split(parts[1], ", ")

		for _, edge := range edges {
			matches := edgeMatcher.FindStringSubmatch(edge)
			if len(matches) > 0 {
				color := matches[1]
				if _, ok := g.nodes[color]; !ok {
					g.addNode(color)
				}
				g.addEdge(n, color)
			}
		}
	}
	return g
}
