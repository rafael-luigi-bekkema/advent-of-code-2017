package main

import "testing"

func TestDay5(t *testing.T) {
	TestEqual(t, 5, day5a([]int{0, 3, 0, 1, -3}))
	TestEqual(t, 10, day5b([]int{0, 3, 0, 1, -3}))

	ints := IntLines(5)
	TestEqual(t, 388611, day5a(ints))
	TestEqual(t, 27763113, day5b(ints))
}
