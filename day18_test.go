package main

import "testing"

func TestDay18a(t *testing.T) {
	example := []string{
		"set a 1",
		"add a 2",
		"mul a a",
		"mod a 5",
		"snd a",
		"set a 0",
		"rcv a",
		"jgz a -1",
		"set a 1",
		"jgz a -2",
	}
	TestEqual(t, 4, day18a(example))

	exampleb := []string{
		"snd 1",
		"snd 2",
		"snd p",
		"rcv a",
		"rcv b",
		"rcv c",
		"rcv d",
	}
	TestEqual(t, 3, day18b(exampleb))

	file := Lines(18)
	TestEqual(t, 8600, day18a(file))
	TestEqual(t, 7239, day18b(file))
}
