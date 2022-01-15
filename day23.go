package main

import (
	"fmt"
	"strings"
)

type CPU struct {
	regs     [8]int
	mulCount int
	ptr      int
	cmds     []func()
}

func (cpu *CPU) Get(s string) *int {
	if 'a' <= s[0] && s[0] <= 'z' {
		return &cpu.regs[s[0]-'a']
	}
	i := Atoi(s)
	return &i
}

func (cpu *CPU) Set(s string, val int) {
	if len(s) > 1 {
		panic("invalid register: " + s)
	}
	cpu.regs[s[0]-'a'] = val
}

/*
set f 1
set d 2
set e 2

set g d
mul g e
sub g b
jnz g 2
set f 0
sub e -1
set g e
sub g b
jnz g -8

if d*e - b == 0 {
	f = 0
}
e++
if e - b != 0 { repeat }
*/

func (cpu *CPU) Load(cmds []string, optimal bool) {
	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]
		// out := fmt.Sprintln(cmd, "ptr:", i)
		// fmt.Println(cmd)

		parts := strings.Split(cmd, " ")
		vv1, vv2 := cpu.Get(parts[1]), cpu.Get(parts[2])
		switch parts[0] {
		case "set":
			cpu.cmds = append(cpu.cmds, func() {
				if cpu.ptr == 11 && optimal {
					d := cpu.regs['b'-'a']
					b := cpu.regs['b'-'a']
					e := cpu.regs['e'-'a']
					// cpu.regs['f'-'a'] = cpu.regs['b'-'a'] % cpu.regs['d'-'a']
					fmt.Println(d, e, b)
					for e := 2; e <= b; e++ {
						if d*e == b {
							cpu.regs['f'-'a'] = 0
							break
						}
					}
					cpu.regs['e'-'a'] = cpu.regs['b'-'a']
					cpu.ptr += 8
					return
					// fmt.Println("sub", val, cpu.regs['b'-'a'])
				}
				*vv1 = *vv2
			})
		case "sub":
			cpu.cmds = append(cpu.cmds, func() {
				*vv1 -= *vv2
			})
		case "mul":
			cpu.cmds = append(cpu.cmds, func() {
				cpu.mulCount++
				*vv1 *= *vv2
			})
		case "jnz":
			cpu.cmds = append(cpu.cmds, func() {
				if *vv1 != 0 {
					cpu.ptr += *vv2 - 1
					if cpu.ptr+1 >= len(cmds) {
						fmt.Println("end", cmd)
					}
				}
			})
		default:
			panic("unknown cmd: " + cmd)
		}
	}
}

func (cpu *CPU) Run() {
	for cpu.ptr = 0; cpu.ptr < len(cpu.cmds); cpu.ptr++ {
		cpu.cmds[cpu.ptr]()
	}
}

func day23a(cmds []string) int {
	cpu := CPU{}
	cpu.Load(cmds, false)
	cpu.Run()
	return cpu.mulCount
}

func day23b(cmds []string) int {
	// cpu := CPU{}
	// cpu.regs[0] = 1
	// cpu.Load(cmds, true)
	// cpu.Run()
	// return cpu.regs['h'-'a']

	// Rewritten and simplified
	b := 79*100 + 100000
	c := b + 17000
	h := 0
	for ; b <= c; b += 17 {
		for d := 2; d < b; d++ {
			if b%d == 0 {
				h++
				break
			}
		}
	}
	return h
}
