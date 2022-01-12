package main

import (
	"fmt"
	"strings"
)

type day7node struct {
	name          string
	weight, total int
	parent        *day7node
	children      []*day7node
}

func (n *day7node) tweight(depth int) (int, int) {
	weight := n.weight

	weights := make([]int, len(n.children))
	weightMap := map[int]int{}
	childMap := map[int]*day7node{}
	balanced := true
	var found int
	for i, cn := range n.children {
		var snd int
		weights[i], snd = cn.tweight(depth + 1)
		if snd != 0 {
			found = snd
		}
		if i > 0 && weights[i] != weights[0] {
			balanced = false
		}
		weightMap[weights[i]]++
		childMap[weights[i]] = cn
		weight += weights[i]
	}
	if !balanced && found == 0 {
		for weight, count := range weightMap {
			if count == 1 {
				for weight2, count := range weightMap {
					if count > 1 {
						cweight := childMap[weight].weight + (weight2 - weight)
						return weight, cweight
					}
				}
			}
		}
	}
	return weight, found
}

func day7nodes(input []string) map[string]*day7node {
	nodes := make(map[string]*day7node, len(input))
	for _, line := range input {
		p, cs, _ := strings.Cut(line, " -> ")
		var node day7node
		fmt.Sscanf(p, "%s (%d)", &node.name, &node.weight)
		if nodes[node.name] == nil {
			nodes[node.name] = &node
		} else {
			nodes[node.name].weight = node.weight
		}

		for _, c := range strings.Split(cs, ", ") {
			if nodes[c] == nil {
				nodes[c] = &day7node{name: c}
			}
			nodes[c].parent = &node
			nodes[node.name].children = append(nodes[node.name].children, nodes[c])
		}
	}
	return nodes
}

func day7root(nodes map[string]*day7node) *day7node {
	for _, n := range nodes {
		if n.parent == nil {
			return n
		}
	}
	return nil
}

func day7a(input []string) string {
	nodes := day7nodes(input)
	return day7root(nodes).name
}

func day7b(input []string) int {
	nodes := day7nodes(input)
	root := day7root(nodes)
	_, found := root.tweight(1)
	return found
}
