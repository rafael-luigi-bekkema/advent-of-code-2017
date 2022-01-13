package main

import (
	"fmt"
	"strings"
)

func day10a(n int, input string) int {
	nums := Map(Atoi, strings.Split(input, ","))
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = i
	}
	var pos, skip int
	for _, num := range nums {
		for i, j := 0, num-1; i < j; i, j = i+1, j-1 {
			ints[i], ints[j] = ints[j], ints[i]
		}
		mv := (skip + num) % n
		shiftL(ints, mv)
		pos = (pos + (n - mv)) % n
		skip++
	}
	return ints[pos] * ints[(pos+1)%n]
}

func xor(ints []int) int {
	val := ints[0]
	for _, i := range ints[1:] {
		val ^= i
	}
	return val
}

func hexify(input []int) string {
	hex := make([]string, len(input))
	for i, val := range input {
		hex[i] = fmt.Sprintf("%02x", val)
	}
	return strings.Join(hex, "")
}

func shiftL(ints []int, mv int) {
	tmp := make([]int, mv)
	copy(tmp, ints)
	copy(ints, ints[mv:])
	copy(ints[len(ints)-mv:], tmp)
}

func KnotHash(input string) string {
	n := 256
	nums := append([]byte(input), 17, 31, 73, 47, 23)
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = i
	}
	var pos, skip int
	for round := 0; round < 64; round++ {
		for _, num := range nums {
			for i, j := byte(0), num-1; i < j; i, j = i+1, j-1 {
				ints[i], ints[j] = ints[j], ints[i]
			}
			mv := (skip + int(num)) % n
			shiftL(ints, mv)
			pos = (pos + (n - mv)) % n
			skip++
		}
	}

	shiftL(ints, pos)
	dense := make([]int, n/16)
	for i := 0; i < len(dense); i++ {
		dense[i] = xor(ints[i*16 : i*16+16])
	}
	return hexify(dense)
}

func day10b(input string) string {
	return KnotHash(input)
}
