package main

import (
	"container/ring"
)

func day17a(step int) int {
	const repeat = 2017
	cur := ring.New(1)
	cur.Value = 0
	for n := 1; n <= repeat; n++ {
		cur = cur.Move(step - 1)
		newr := ring.New(1)
		newr.Value = n
		cur = cur.Link(newr)
	}
	return cur.Value.(int)
}

func day17b(step int) int {
	const repeat = 50_000_000

	clen := 1
	ptr := 0

	oneval := 1
	oneptr := 1
	for n := 1; n <= repeat; n++ {
		ptr = (ptr + step) % clen
		clen++
		ptr++
		if ptr == oneptr {
			oneval = n
		}
		if ptr < oneptr {
			oneptr++
		}
	}
	return oneval
}
