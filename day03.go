package main

const (
	right = iota
	up
	left
	down
)

type day3 struct {
	x, y, step, remain int
	dir                int
}

func newDay3() *day3 {
	return &day3{step: 1, dir: right, remain: 1}
}

func (d3 *day3) inc() {
	switch d3.dir {
	case right:
		d3.x++
	case up:
		d3.y--
	case left:
		d3.x--
	case down:
		d3.y++
	}
	
	if d3.remain--; d3.remain == 0 {
		d3.step += d3.dir % 2
		d3.dir = (d3.dir + 1) % 4
		d3.remain = d3.step
	}
}

func day3a(input int) int {
	d3 := newDay3()
	for i := 1; i < input; i++ {
		d3.inc()
	}
	return Abs(d3.x) + Abs(d3.y)
}

func day3b(input int) int {
	d3 := newDay3()
	values := map[[2]int]int{{0, 0}: 1}
	for {
		d3.inc()

		point := [2]int{d3.x, d3.y}
		for j := 0; j < 9; j++ {
			if j == 4 {
				continue
			}
			apoint := [2]int{point[0] + (j%3 - 1), point[1] + (j/3 - 1)}
			values[point] += values[apoint]
		}
		if values[point] > input {
			return values[point]
		}
	}
}
