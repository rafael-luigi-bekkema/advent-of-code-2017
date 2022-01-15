package main

import "testing"

func TestDay22(t *testing.T) {
	example := []string{
		"..#",
		"#..",
		"...",
	}
	TestEqual(t, 41, day22a(example, 70))
	TestEqual(t, 5587, day22a(example, 10000))

	TestEqual(t, 26, day22b(example, 100))
	TestEqual(t, 2_511_944, day22b(example, 10_000_000))

	file := Lines(22)
	TestEqual(t, 5369, day22a(file, 10000))
	TestEqual(t, 2510774, day22b(file, 10_000_000))
}
