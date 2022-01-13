package main

import (
	"strings"
)

func day12a(input []string) int {
	links := map[int]map[int]bool{}
	for _, line := range input {
		sfrom, rest, _ := strings.Cut(line, " <-> ")
		from := Atoi(sfrom)
		tos := Map(Atoi, strings.Split(rest, ", "))
		links[from] = make(map[int]bool)
		for _, to := range tos {
			links[from][to] = true
			if links[to] == nil {
				links[to] = make(map[int]bool)
			}
			links[to][from] = true
		}
	}

	var total int
	queue := []int{0}
	queued := map[int]bool{0: true}
	for len(queue) > 0 {
		for l := range links[queue[0]] {
			if queued[l] {
				continue
			}
			queued[l] = true
			queue = append(queue, l)
		}
		queue = queue[1:]
		total++
	}

	return total
}

func day12b(input []string) int {
	links := map[int]map[int]bool{}
	for _, line := range input {
		sfrom, rest, _ := strings.Cut(line, " <-> ")
		from := Atoi(sfrom)
		tos := Map(Atoi, strings.Split(rest, ", "))
		links[from] = make(map[int]bool)
		for _, to := range tos {
			links[from][to] = true
			if links[to] == nil {
				links[to] = make(map[int]bool)
			}
			links[to][from] = true
		}
	}

	var groups int
	for len(links) > 0 {
		groups++
		var first int
		for l := range links {
			first = l
			break
		}

		queue := []int{first}
		queued := map[int]bool{first: true}
		for len(queue) > 0 {
			for l := range links[queue[0]] {
				if queued[l] {
					continue
				}
				queued[l] = true
				queue = append(queue, l)
			}

			delete(links, queue[0])
			queue = queue[1:]
		}
	}

	return groups
}
