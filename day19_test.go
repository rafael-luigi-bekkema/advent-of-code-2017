package main

import "testing"

func TestDay19(t *testing.T) {
	example := "" +
		"     |         \n" +
		"     |  +--+   \n" +
		"     A  |  C   \n" +
		" F---|----E|--+\n" +
		"     |  |  |  D\n" +
		"     +B-+  +--+"
	r1, r2 := day19(example)
	TestEqual(t, "ABCDEF", r1)
	TestEqual(t, 38, r2)

	file := Input(19)
	r1, r2 = day19(file)
	TestEqual(t, "GINOWKYXH", r1)
	TestEqual(t, 16636, r2)
}
