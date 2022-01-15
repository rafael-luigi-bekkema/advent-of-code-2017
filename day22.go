package main

type Pos struct {
	x, y int
}

type Direction Pos

func day22a(input []string, n int) int {
	dirs := []Direction{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	grid := make([][]byte, len(input))
	for i, line := range input {
		grid[i] = []byte(line)
	}
	grids := map[Pos][][]byte{
		{0, 0}: grid,
	}
	gridpos := Pos{0, 0}
	size := len(grid)
	x := size / 2
	y := x
	dir := 0
	var count int
	for burst := 0; burst < n; burst++ {
		if grid[y][x] == '#' {
			dir = (dir + 1) % 4
			grid[y][x] = '.'
		} else {
			dir = (dir + 3) % 4
			grid[y][x] = '#'
			count++
		}
		x, y = x+dirs[dir].x, y+dirs[dir].y
		switch {
		case x < 0:
			gridpos.x--
			x = size - 1
		case x > size-1:
			gridpos.x++
			x = 0
		case y < 0:
			gridpos.y--
			y = size - 1
		case y > size-1:
			gridpos.y++
			y = 0
		default:
			continue
		}
		if _, ok := grids[gridpos]; !ok {
			grids[gridpos] = NewGrid(size, byte('.'))
		}
		grid = grids[gridpos]
	}
	return count
}

func day22b(input []string, n int) int {
	dirs := []Direction{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	grid := make([][]byte, len(input))
	for i, line := range input {
		grid[i] = []byte(line)
	}
	grids := map[Pos][][]byte{
		{0, 0}: grid,
	}
	gridpos := Pos{0, 0}
	size := len(grid)
	x := size / 2
	y := x
	dir := 0
	var count int
	for burst := 0; burst < n; burst++ {
		switch grid[y][x] {
		case '#':
			grid[y][x] = 'F'
			dir = (dir + 1) % 4
		case '.':
			grid[y][x] = 'W'
			dir = (dir + 3) % 4
		case 'F':
			grid[y][x] = '.'
			dir = (dir + 2) % 4
		case 'W':
			grid[y][x] = '#'
			count++
		}
		x, y = x+dirs[dir].x, y+dirs[dir].y
		switch {
		case x < 0:
			gridpos.x--
			x = size - 1
		case x > size-1:
			gridpos.x++
			x = 0
		case y < 0:
			gridpos.y--
			y = size - 1
		case y > size-1:
			gridpos.y++
			y = 0
		default:
			continue
		}
		if _, ok := grids[gridpos]; !ok {
			grids[gridpos] = NewGrid(size, byte('.'))
		}
		grid = grids[gridpos]
	}
	return count
}
