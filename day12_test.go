package main

import "testing"

func TestDay12(t *testing.T) {
	input := []string{
		"0 <-> 2",
		"1 <-> 1",
		"2 <-> 0, 3, 4",
		"3 <-> 2, 4",
		"4 <-> 2, 3, 6",
		"5 <-> 6",
		"6 <-> 4, 5",
	}
	TestEqual(t, 6, day12a(input))
	TestEqual(t, 2, day12b(input))

	file := Lines(12)
	TestEqual(t, 306, day12a(file))
	TestEqual(t, 200, day12b(file))
}
