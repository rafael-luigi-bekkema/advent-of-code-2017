package main

import "testing"

func TestDay14(t *testing.T) {
	TestEqual(t, 8108, day14a("flqrgnkx"))
	TestEqual(t, 1242, day14b("flqrgnkx"))
	TestEqual(t, 8230, day14a("hfdlxzhv"))
	TestEqual(t, 1103, day14b("hfdlxzhv"))
}
