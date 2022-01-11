package main

import "testing"

func TestDay3(t *testing.T) {
	TestEqual(t, 0, day3a(1))
	TestEqual(t, 3, day3a(12))
	TestEqual(t, 2, day3a(23))
	TestEqual(t, 31, day3a(1024))
	TestEqual(t, 475, day3a(277678))

	TestEqual(t, 5, day3b(4))
	TestEqual(t, 57, day3b(56))
	TestEqual(t, 747, day3b(700))
	TestEqual(t, 279138, day3b(277678))
}
