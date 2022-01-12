package main

import "fmt"

func day6(mem []int) (steps int) {
	hist := map[string]bool{
		fmt.Sprint(mem): true,
	}
	for {
		steps++
		var maxi, maxq int
		for i := range mem {
			if mem[i] > maxq {
				maxq = mem[i]
				maxi = i
			}
		}

		mem[maxi] = 0
		for i := 1; i <= maxq; i++ {
			mem[(maxi+i)%len(mem)]++
		}

		s := fmt.Sprint(mem)
		if _, ok := hist[s]; ok {
			break
		}
		hist[s] = true
	}
	return
}

func day6a(mem ...int) (steps int) {
	return day6(CopySlice(mem))
}

func day6b(mem ...int) (loop int) {
	mem = CopySlice(mem)
	day6(mem)
	res := day6(mem)
	return res
}
