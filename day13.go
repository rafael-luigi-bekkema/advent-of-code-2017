package main

import (
	"fmt"
)

func day13a(input []string) int {
	scanners := map[int]int{}
	var depth int
	for _, line := range input {
		var rng int
		fmt.Sscanf(line, "%d: %d", &depth, &rng)
		scanners[depth] = rng

	}
	var severity int
	for s := 0; s <= depth; s++ {
		if rng, ok := scanners[s]; ok {
			cycle := (rng - 1) * 2
			if s%cycle == 0 {
				severity += rng * s
			}
		}
	}
	return severity
}

func day13b(input []string) int {
	scanners := map[int]int{}
	var depth int
	for _, line := range input {
		var rng int
		fmt.Sscanf(line, "%d: %d", &depth, &rng)
		scanners[depth] = rng

	}
	var delay int
delay:
	for {
		for s := 0; s <= depth; s++ {
			if rng, ok := scanners[s]; ok {
				cycle := (rng - 1) * 2
				if (s+delay)%cycle == 0 {
					delay++
					continue delay
				}
			}
		}
		break
	}
	return delay
}
