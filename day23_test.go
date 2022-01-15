package main

import "testing"

func TestDay23(t *testing.T) {
	TestEqual(t, 5929, day23a(Lines(23)))
	TestEqual(t, 907, day23b(Lines(23)))
}
