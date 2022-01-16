package main

import "testing"

func TestDay25(t *testing.T) {
	example := []string{
		"Begin in state A.",
		"Perform a diagnostic checksum after 6 steps.",
		"",
		"In state A:",
		"  If the current value is 0:",
		"    - Write the value 1.",
		"    - Move one slot to the right.",
		"    - Continue with state B.",
		"  If the current value is 1:",
		"    - Write the value 0.",
		"    - Move one slot to the left.",
		"    - Continue with state B.",
		"",
		"In state B:",
		"  If the current value is 0:",
		"    - Write the value 1.",
		"    - Move one slot to the left.",
		"    - Continue with state A.",
		"  If the current value is 1:",
		"    - Write the value 1.",
		"    - Move one slot to the right.",
		"    - Continue with state A.",
	}
	TestEqual(t, 3, day25a(example))
	file := Lines(25)
	TestEqual(t, 2870, day25a(file))
}
