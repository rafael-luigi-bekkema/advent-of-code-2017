package main

import "fmt"

func day8a(lines []string) (int, int) {
	regs := make(map[string]int)
	var maxAll int
	for i, line := range lines {
		var reg1, reg2, ins, cmp string
		var v1, v2 int
		Must(fmt.Sscanf(line, "%s %s %d if %s %s %d", &reg1, &ins, &v1, &reg2, &cmp, &v2))
		var cond bool
		switch cmp {
		case ">":
			cond = regs[reg2] > v2
		case "<":
			cond = regs[reg2] < v2
		case ">=":
			cond = regs[reg2] >= v2
		case "<=":
			cond = regs[reg2] <= v2
		case "==":
			cond = regs[reg2] == v2
		case "!=":
			cond = regs[reg2] != v2
		default:
			panic("unknown operator: " + cmp)
		}

		if cond {
			if ins == "dec" {
				v1 *= -1
			}
			regs[reg1] += v1
			if i == 0 || regs[reg1] > maxAll {
				maxAll = regs[reg1]
			}
		}
	}
	_, max := MinMax(Values(regs))
	return max, maxAll
}
