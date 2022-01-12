package main

import "testing"

func TestDay8(t *testing.T) {
	example := []string{
		"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10",
	}
	v1, v2 := day8a(example)
	TestEqual(t, 1, v1)
	TestEqual(t, 10, v2)

	v1, v2 = day8a(Lines(8))
	TestEqual(t, 4888, v1)
	TestEqual(t, 7774, v2)
}
