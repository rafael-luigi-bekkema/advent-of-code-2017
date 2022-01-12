package main

import "testing"

func TestDay7(t *testing.T) {
	TestEqual(t, "tknk", day7a([]string{
		"pbga (66)",
		"xhth (57)",
		"ebii (61)",
		"havc (66)",
		"ktlj (57)",
		"fwft (72) -> ktlj, cntj, xhth",
		"qoyq (66)",
		"padx (45) -> pbga, havc, qoyq",
		"tknk (41) -> ugml, padx, fwft",
		"jptl (61)",
		"ugml (68) -> gyxo, ebii, jptl",
		"gyxo (61)",
		"cntj (57)",
	}))
	TestEqual(t, 60, day7b([]string{
		"pbga (66)",
		"xhth (57)",
		"ebii (61)",
		"havc (66)",
		"ktlj (57)",
		"fwft (72) -> ktlj, cntj, xhth",
		"qoyq (66)",
		"padx (45) -> pbga, havc, qoyq",
		"tknk (41) -> ugml, padx, fwft",
		"jptl (61)",
		"ugml (68) -> gyxo, ebii, jptl",
		"gyxo (61)",
		"cntj (57)",
	}))
	file := Lines(7)

	TestEqual(t, "eugwuhl", day7a(file))
	TestEqual(t, 420, day7b(file))
}
