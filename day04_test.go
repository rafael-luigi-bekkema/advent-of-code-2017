package main

import "testing"

func TestDay4(t *testing.T) {
	TestEqual(t, 2, day4a([]string{
		"aa bb cc dd ee",
		"aa bb cc dd aa",
		"aa bb cc dd aaa",
	}))

	TestEqual(t, 3, day4b([]string{
		"abcde fghij",
		"abcde xyz ecdab",
		"a ab abc abd abf abj",
		"iiii oiii ooii oooi oooo",
		"oiii ioii iioi iiio is",
	}))

	file := Lines(4)
	TestEqual(t, 383, day4a(file))
	TestEqual(t, 265, day4b(file))
}
