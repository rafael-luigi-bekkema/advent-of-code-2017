package main

import "testing"

func TestDay2(t *testing.T) {
	TestEqual(t, 18, day2a([]string{"5 1 9 5",
		"7 5 3",
		"2 4 6 8"}))

	file := Lines(2)
	TestEqual(t, 53460, day2a(file))

	TestEqual(t, 9, day2b([]string{"5 9 2 8",
		"9 4 7 3",
		"3 8 6 5"}))
	TestEqual(t, 282, day2b(file))
}
