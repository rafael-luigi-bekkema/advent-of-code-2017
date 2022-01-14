package main

import (
	"strings"
	"sync"
)

func day18a(cmds []string) int {
	regs := map[byte]int{}
	get := func(s string) int {
		if 'a' <= s[0] && s[0] <= 'z' {
			return regs[s[0]]
		}
		return Atoi(s)
	}
	set := func(s string, val int) {
		if len(s) > 1 {
			panic("invalid register: " + s)
		}
		regs[s[0]] = val
	}
	var lastSnd int
	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]
		ins, rest, _ := strings.Cut(cmd, " ")
		switch ins {
		case "snd":
			lastSnd = get(rest)
		case "set":
			v1, v2, _ := strings.Cut(rest, " ")
			set(v1, get(v2))
		case "add":
			v1, v2, _ := strings.Cut(rest, " ")
			set(v1, get(v1)+get(v2))
		case "mul":
			v1, v2, _ := strings.Cut(rest, " ")
			set(v1, get(v1)*get(v2))
		case "mod":
			v1, v2, _ := strings.Cut(rest, " ")
			set(v1, get(v1)%get(v2))
		case "rcv":
			if get(rest) != 0 {
				return lastSnd
			}
		case "jgz":
			v1, v2, _ := strings.Cut(rest, " ")
			if get(v1) > 0 {
				i += get(v2) - 1
			}
		default:
			panic("unknown cmd: " + cmd)
		}
	}
	return 0
}

func day18b(cmds []string) int {
	var lock sync.Mutex
	var locks int

	var rcount int

	var wg sync.WaitGroup
	program := func(pid int, snd, rcv chan int) {
		go func() {
			defer wg.Done()
			regs := map[byte]int{
				'p': pid,
			}
			get := func(s string) int {
				if 'a' <= s[0] && s[0] <= 'z' {
					return regs[s[0]]
				}
				return Atoi(s)
			}
			set := func(s string, val int) {
				if len(s) > 1 {
					panic("invalid register: " + s)
				}
				regs[s[0]] = val
			}
			var count int
			for i := 0; i < len(cmds); i++ {
				cmd := cmds[i]
				ins, rest, _ := strings.Cut(cmd, " ")
				switch ins {
				case "snd":
					snd <- get(rest)
					count++
				case "set":
					v1, v2, _ := strings.Cut(rest, " ")
					set(v1, get(v2))
				case "add":
					v1, v2, _ := strings.Cut(rest, " ")
					set(v1, get(v1)+get(v2))
				case "mul":
					v1, v2, _ := strings.Cut(rest, " ")
					set(v1, get(v1)*get(v2))
				case "mod":
					v1, v2, _ := strings.Cut(rest, " ")
					set(v1, get(v1)%get(v2))
				case "rcv":
					for len(rcv) == 0 {
						lock.Lock()
						locks++
						lock.Unlock()
						lock.Lock()
						// This can only be true if the other program is waiting
						// to get into this block and since both channels are empty
						// we have a deadlock
						if locks >= 2 && len(snd) == 0 && len(rcv) == 0 {
							lock.Unlock()
							if pid == 1 {
								rcount = count
							}
							return
						}
						locks--
						lock.Unlock()
					}
					set(rest, <-rcv)
				case "jgz":
					v1, v2, _ := strings.Cut(rest, " ")
					if get(v1) > 0 {
						i += get(v2) - 1
					}
				default:
					panic("unknown cmd: " + cmd)
				}
			}
		}()
	}

	c0, c1 := make(chan int, 100), make(chan int, 100)
	wg.Add(2)
	program(0, c0, c1)
	program(1, c1, c0)
	wg.Wait()
	return rcount
}
