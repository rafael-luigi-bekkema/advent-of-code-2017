package main

import "testing"

func TestDay1(t *testing.T) {
	TestEqual(t, 3, day1a("1122"))
	TestEqual(t, 4, day1a("1111"))
	TestEqual(t, 0, day1a("1234"))
	TestEqual(t, 9, day1a("91212129"))

	TestEqual(t, 6, day1b("1212"))
	TestEqual(t, 0, day1b("1221"))
	TestEqual(t, 4, day1b("123425"))
	TestEqual(t, 12, day1b("123123"))
	TestEqual(t, 4, day1b("12131415"))

	file := Input(1)
	TestEqual(t, 1175, day1a(file))
	TestEqual(t, 1166, day1b(file))
}
