package main

import "fmt"

func day14a(key string) int {
	var active int
	for rowi := 0; rowi < 128; rowi++ {
		hash := KnotHash(fmt.Sprintf("%s-%d", key, rowi))
		for _, chr := range []byte(hash) {
			var val byte
			if '0' <= chr && chr <= '9' {
				val = chr - '0'
			} else {
				val = (chr - 'a') + 10
			}
			for _, hchr := range []byte(fmt.Sprintf("%04b", val)) {
				if hchr == '1' {
					active++
				}
			}
		}
	}
	return active
}

func day14b(key string) int {
	grid := map[[2]int]bool{}
	for y := 0; y < 128; y++ {
		hash := KnotHash(fmt.Sprintf("%s-%d", key, y))
		for i, chr := range []byte(hash) {
			var val byte
			if '0' <= chr && chr <= '9' {
				val = chr - '0'
			} else {
				val = (chr - 'a') + 10
			}
			for j, hchr := range []byte(fmt.Sprintf("%04b", val)) {
				if hchr == '1' {
					grid[[2]int{i*4 + j, y}] = true
				}
			}
		}
	}
	var zones int
	for len(grid) > 0 {
		var first [2]int
		for c := range grid {
			first = c
			break
		}
		queue := [][2]int{first}
		queued := map[[2]int]bool{}
		for len(queue) > 0 {
			queued[queue[0]] = true
			x, y := queue[0][0], queue[0][1]
			for _, off := range [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}} {
				dc := [2]int{x + off[0], y + off[1]}
				if queued[dc] {
					continue
				}
				if grid[dc] {
					queue = append(queue, dc)
					queued[dc] = true
				}
			}
			delete(grid, queue[0])
			queue = queue[1:]
		}
		zones++
	}
	return zones
}
