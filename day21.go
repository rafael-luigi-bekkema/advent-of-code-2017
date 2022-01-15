package main

import (
	"fmt"
	"strings"
)

func newG(size int) [][]byte {
	newg := make([][]byte, size)
	for i := range newg {
		newg[i] = make([]byte, size)
	}
	return newg
}

func rotate(grid [][]byte) [][]byte {
	size := len(grid)
	newg := newG(size)
	for y, row := range grid {
		for x, chr := range row {
			newg[x][size-y-1] = chr
		}
	}
	return newg
}

func flipV(grid [][]byte) [][]byte {
	grid = copyG(grid)
	for x := 0; x < len(grid[0]); x++ {
		for i, j := 0, len(grid)-1; i < j; i, j = i+1, j-1 {
			grid[i][x], grid[j][x] = grid[j][x], grid[i][x]
		}
	}
	return grid
}

func flipH(grid [][]byte) [][]byte {
	grid = copyG(grid)
	for y := range grid {
		for i, j := 0, len(grid[y])-1; i < j; i, j = i+1, j-1 {
			grid[y][i], grid[y][j] = grid[y][j], grid[y][i]
		}
	}
	return grid
}

func copyG(grid [][]byte) [][]byte {
	size := len(grid)
	newg := make([][]byte, size)
	for y, row := range grid {
		newg[y] = make([]byte, size)
		copy(newg[y], row)
	}
	return newg
}

func fromS(input string) [][]byte {
	var size int
	switch len(input) {
	case 4:
		size = 2
	case 9:
		size = 3
	case 16:
		size = 4
	default:
		panic("uknown size: " + input)
	}
	newg := newG(size)
	for i, c := range []byte(input) {
		newg[i/size][i%size] = c
	}
	return newg
}

func toS(grid [][]byte) string {
	size := len(grid)
	s := make([]byte, size*size)
	for y, row := range grid {
		for x, c := range row {
			s[y*size+x] = c
		}
	}
	return string(s)
}

func render(grid [][]byte) {
	fmt.Println()
	for _, row := range grid {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func day21a(lines []string, n int) int {
	grid := [][]byte{
		{'.', '#', '.'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	}
	pat2 := map[string]string{}
	pat3 := map[string]string{}
	add := func(from, to string) {
		// render(fromS(from))
		if len(from) == 4 {
			pat2[from] = to
		} else {
			pat3[from] = to
		}
	}
	for _, line := range lines {
		from, to, _ := strings.Cut(strings.ReplaceAll(line, "/", ""), " => ")
		fromG := fromS(from)
		for i := 0; i < 4; i++ {
			add(from, to)
			add(toS(flipH(fromG)), to)
			add(toS(flipV(fromG)), to)
			if i < 3 {
				fromG = rotate(fromG)
				from = toS(fromG)
			}
		}
	}
	for i := 0; i < n; i++ {
		size := len(grid)

		if size%2 == 0 {
			newsize := (size / 2) * 3
			newg := newG(newsize)
			for y := 0; y < size; y += 2 {
				for x := 0; x < size; x += 2 {
					p := string([]byte{grid[y][x], grid[y][x+1],
						grid[y+1][x], grid[y+1][x+1]})
					if to, ok := pat2[p]; ok {
						toG := fromS(to)
						dy, dx := (y/2)*3, (x/2)*3
						for ny, nrow := range toG {
							for nx, nc := range nrow {
								newg[dy+ny][dx+nx] = nc
							}
						}
					}
				}
			}
			grid = newg
			continue
		}

		newsize := (size / 3) * 4
		newg := newG(newsize)
		for y := 0; y < size; y += 3 {
			for x := 0; x < size; x += 3 {
				p := string([]byte{grid[y][x], grid[y][x+1], grid[y][x+2],
					grid[y+1][x], grid[y+1][x+1], grid[y+1][x+2],
					grid[y+2][x], grid[y+2][x+1], grid[y+2][x+2]})
				if to, ok := pat3[p]; ok {
					toG := fromS(to)
					dy, dx := (y/3)*4, (x/3)*4
					for ny, nrow := range toG {
						for nx, nc := range nrow {
							newg[dy+ny][dx+nx] = nc
						}
					}
				}
			}
		}
		grid = newg
	}
	var count int
	for _, row := range grid {
		for _, c := range row {
			if c == '#' {
				count++
			}
		}
	}
	return count
}
