package main

import (
	"testing"
)

func TestDay17(t *testing.T) {
	TestEqual(t, 638, day17a(3))
	TestEqual(t, 1025, day17a(366))
	TestEqual(t, 37803463, day17b(366))
}
