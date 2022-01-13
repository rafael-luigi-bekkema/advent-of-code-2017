package main

import (
	"strings"
)

func day11b(input string) int {
	_, result := day11(input)
	return result
}

func day11a(input string) int {
	result, _ := day11(input)
	return result
}

func day11(input string) (int, int) {
	var q, r int
	dist := func() int {
		s := -q - r
		return (Abs(q) + Abs(r) + Abs(s)) / 2
	}
	var maxdist int
	for _, step := range strings.Split(input, ",") {
		switch step {
		case "n":
			r--
		case "nw":
			q--
		case "sw":
			q--
			r++
		case "s":
			r++
		case "se":
			q++
		case "ne":
			q++
			r--
		}
		if d := dist(); d > maxdist {
			maxdist = d
		}
	}
	return dist(), maxdist
}
