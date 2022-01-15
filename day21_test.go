package main

import "testing"

func TestDay21(t *testing.T) {
	example := []string{
		"../.# => ##./#../...",
		".#./..#/### => #..#/..../..../#..#",
	}
	TestEqual(t, 12, day21a(example, 2))

	file := Lines(21)
	TestEqual(t, 123, day21a(file, 5))
	TestEqual(t, 1984683, day21a(file, 18))
}
