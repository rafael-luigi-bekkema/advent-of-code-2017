package main

import (
	"strings"
	"testing"
)

func TestDay6(t *testing.T) {
	TestEqual(t, 5, day6a(0, 2, 7, 0))
	TestEqual(t, 4, day6b(0, 2, 7, 0))

	file := Map(Atoi, strings.Fields(Input(6)))
	TestEqual(t, 11137, day6a(file...))
	TestEqual(t, 1037, day6b(file...))
}
