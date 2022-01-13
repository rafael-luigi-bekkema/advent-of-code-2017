package main

import "fmt"

const day15div = 2_147_483_647

func day15lastBit(val int) string {
	s := fmt.Sprintf("%16b", val)
	return s[len(s)-16:]
}

func day15aStartGen(prev, factor int) chan string {
	gen := make(chan string, 50)
	go func() {
		for {
			prev = (prev * factor) % day15div
			gen <- day15lastBit(prev)
		}
	}()
	return gen
}

func day15a(prevA, prevB, n int) int {
	var total int
	factorA, factorB := 16807, 48271
	genA, genB := day15aStartGen(prevA, factorA), day15aStartGen(prevB, factorB)
	for i := 0; i < n; i++ {
		if <-genA == <-genB {
			total++
		}
	}
	return total
}

func day15bStartGen(prev, factor, mod int) chan string {
	gen := make(chan string, 50)
	go func() {
		for {
			prev = (prev * factor) % day15div
			if prev%mod == 0 {
				gen <- day15lastBit(prev)
			}
		}
	}()
	return gen
}

func day15b(prevA, prevB, n int) int {
	var total int
	factorA, factorB := 16807, 48271
	genA := day15bStartGen(prevA, factorA, 4)
	genB := day15bStartGen(prevB, factorB, 8)

	for i := 0; i < n; i++ {
		valA := <-genA
		valB := <-genB

		if valA == valB {
			total++
		}
	}
	return total
}
