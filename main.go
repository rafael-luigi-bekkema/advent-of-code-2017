package main

import "fmt"

func main() {
	input := []string{
		"0 <-> 2",
		"1 <-> 1",
		"2 <-> 0, 3, 4",
		"3 <-> 2, 4",
		"4 <-> 2, 3, 6",
		"5 <-> 6",
		"6 <-> 4, 5",
	}
	fmt.Println("day 12a:", day12a(input))
}
