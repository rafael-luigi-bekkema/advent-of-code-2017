package main

import "testing"

func TestDay20(t *testing.T) {
	TestEqual(t, 0, day20a([]string{
		"p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>",
		"p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>",
	}))
	TestEqual(t, 1, day20b([]string{
		"p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>",
		"p=<-4,0,0>, v=<2,0,0>, a=<0,0,0>",
		"p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>",
		"p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>",
	}))
	file := Lines(20)
	TestEqual(t, 91, day20a(file))
	TestEqual(t, 567, day20b(file))
}
