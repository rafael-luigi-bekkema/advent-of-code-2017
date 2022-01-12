package main

import (
	"testing"
)

func TestDay10(t *testing.T) {
	TestEqual(t, 12, day10a(5, "3,4,1,5"))

	TestEqual(t, "4007ff", hexify([]int{64, 7, 255}))

	TestEqual(t, "63960835bcdc130f0b66d7ff4f6a5a8e", day10b("1,2,4"))
	TestEqual(t, "a2582a3a0e66e6e86e3812dcb672a272", day10b(""))
	TestEqual(t, "33efeb34ea91902bb2f59c9920caa6cd", day10b("AoC 2017"))
	TestEqual(t, "3efbe78a8d82f29979031a4aa0b16a9d", day10b("1,2,3"))
	TestEqual(t, "63960835bcdc130f0b66d7ff4f6a5a8e", day10b("1,2,4"))

	file := Input(10)
	TestEqual(t, 13760, day10a(256, file))
	TestEqual(t, "2da93395f1a6bb3472203252e3b17fe5", day10b(file))
}
