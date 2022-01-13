package main

import "testing"

func TestDay13(t *testing.T) {
	TestEqual(t, 24, day13a([]string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}))
	TestEqual(t, 10, day13b([]string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}))
	TestEqual(t, 1640, day13a(Lines(13)))
	TestEqual(t, 3960702, day13b(Lines(13)))
}
