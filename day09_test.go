package main

import "testing"

func TestDay9(t *testing.T) {
	TestEqual(t, 1, day9a("{}"))
	TestEqual(t, 6, day9a("{{{}}}"))
	TestEqual(t, 5, day9a("{{},{}}"))
	TestEqual(t, 16, day9a("{{{},{},{{}}}}"))
	TestEqual(t, 1, day9a("{<a>,<a>,<a>,<a>}"))
	TestEqual(t, 9, day9a("{{<ab>},{<ab>},{<ab>},{<ab>}}"))
	TestEqual(t, 9, day9a("{{<!!>},{<!!>},{<!!>},{<!!>}}"))
	TestEqual(t, 3, day9a("{{<a!>},{<a!>},{<a!>},{<ab>}}"))

	TestEqual(t, 0, day9b("<>"))
	TestEqual(t, 17, day9b("<random characters>"))
	TestEqual(t, 3, day9b("<<<<>"))
	TestEqual(t, 2, day9b("<{!>}>"))
	TestEqual(t, 0, day9b("<!!>"))
	TestEqual(t, 0, day9b("<!!!>>"))
	TestEqual(t, 10, day9b(`<{o"i!a,<{i<a>`))

	file := Input(9)
	TestEqual(t, 10800, day9a(file))
	TestEqual(t, 4522, day9b(file))
}
