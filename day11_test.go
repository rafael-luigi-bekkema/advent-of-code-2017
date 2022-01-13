package main

import "testing"

func TestDay11(t *testing.T) {
	TestEqual(t, 3, day11a("ne,ne,ne"))
	TestEqual(t, 0, day11a("ne,ne,sw,sw"))
	TestEqual(t, 2, day11a("ne,ne,s,s"))
	TestEqual(t, 3, day11a("se,sw,se,sw,sw"))

	TestEqual(t, 687, day11a(Input(11)))
	TestEqual(t, 1483, day11b(Input(11)))
}
