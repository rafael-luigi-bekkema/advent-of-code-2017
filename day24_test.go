package main

import "testing"

func TestDay24(t *testing.T) {
	example := []string{
		"0/2",
		"2/2",
		"2/3",
		"3/4",
		"3/5",
		"0/1",
		"10/1",
		"9/10",
	}
	TestEqual(t, 31, day24a(example))
	TestEqual(t, 19, day24b(example))
	file := Lines(24)
	TestEqual(t, 2006, day24a(file))
	TestEqual(t, 1994, day24b(file))
}
