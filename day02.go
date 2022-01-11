package main

import (
	"strings"
)

func day2a(lines []string) (result int) {
	for _, line := range lines {
		nums := Map(Atoi, strings.Fields(line))
		min, max := MinMax(nums)
		result += max - min
	}
	return
}

func day2b(lines []string) (result int) {
	for _, line := range lines {
		nums := Map(Atoi, strings.Fields(line))
		for i, num := range nums {
			for j, num2 := range nums {
				if i == j {
					continue
				}
				if num%num2 == 0 {
					result += num / num2
				}
			}
		}
	}
	return
}
