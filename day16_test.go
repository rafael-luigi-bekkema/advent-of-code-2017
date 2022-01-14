package main

import "testing"

func TestDay16(t *testing.T) {
	TestEqual(t, "baedc", day16a(5, "s1,x3/4,pe/b"))
	file := Input(16)
	TestEqual(t, "giadhmkpcnbfjelo", day16a(16, file))
	TestEqual(t, "njfgilbkcoemhpad", day16b(16, file))
}
