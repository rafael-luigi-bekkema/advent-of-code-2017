package main

import "testing"

func TestDay15a(t *testing.T) {
	if testing.Short() {
		t.Skip("slow test")
	}
	TestEqual(t, 1, day15a(65, 8921, 5))
	TestEqual(t, 588, day15a(65, 8921, 40_000_000))
	TestEqual(t, 638, day15a(289, 629, 40_000_000))
}

func TestDay15b(t *testing.T) {
	if testing.Short() {
		t.Skip("slow test")
	}
	TestEqual(t, 309, day15b(65, 8921, 5_000_000))
	TestEqual(t, 343, day15b(289, 629, 5_000_000))
}
