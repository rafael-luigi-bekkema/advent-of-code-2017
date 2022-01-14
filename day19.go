package main

import (
	"bytes"
	"fmt"
)

func day19(themap string) (string, int) {
	var steps int
	grid := bytes.Split([]byte(themap), []byte("\n"))
	y, x := 0, bytes.Index(grid[0], []byte("|"))
	dy, dx := 1, 0
	var collect []byte
outer:
	for {
		steps++
		x, y = x+dx, y+dy
		switch chr := grid[y][x]; chr {
		case '+':
			if dy != 0 {
				if x < len(grid[y])-1 && grid[y][x+1] != ' ' {
					dy, dx = 0, 1
				} else if x > 0 && grid[y][x-1] != ' ' {
					dy, dx = 0, -1
				}
			} else {
				if y < len(grid)-1 && grid[y+1][x] != ' ' {
					dy, dx = 1, 0
				} else if y > 0 && grid[y-1][x] != ' ' {
					dy, dx = -1, 0
				}
			}
		case '|', '-': // Keep going in same direction
		case ' ':
			break outer
		default:
			if 'A' <= chr && chr <= 'Z' {
				collect = append(collect, chr)
				continue
			}
			panic(fmt.Sprintf("reached invalid location at %dx%d: %q", x, y, chr))

		}
	}
	return string(collect), steps
}
