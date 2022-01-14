package main

import (
	"bytes"
	"fmt"
	"strings"
)

func day16a(n byte, input string) string {
	programs := make([]byte, n)
	for i := byte(0); i < n; i++ {
		programs[i] = i + 'a'
	}
	cmds := strings.Split(input, ",")
	return day16(programs, cmds)
}

func day16b(n byte, input string) string {
	programs := make([]byte, n)
	for i := byte(0); i < n; i++ {
		programs[i] = i + 'a'
	}
	cmds := strings.Split(input, ",")
	const repeats = 1_000_000_000
	hist := map[string]int{}
	for i := 0; i < repeats; i++ {
		day16(programs, cmds)
		s := string(programs)
		// pattern repeats itself after first history match
		// skip to the last cycle
		if _, ok := hist[s]; ok {
			i = repeats - (repeats % i)
			hist = map[string]int{}
		} else {
			hist[s] = i
		}
	}
	return string(programs)
}

func day16(programs []byte, cmds []string) string {
	n := byte(len(programs))
	for _, cmd := range cmds {
		switch cmd[0] {
		case 's':
			cnt := byte(Atoi(cmd[1:]))
			tmp := make([]byte, cnt)
			copy(tmp, programs[n-cnt:])
			copy(programs[cnt:], programs)
			copy(programs, tmp)
		case 'x':
			var pos1, pos2 int
			fmt.Sscanf(cmd[1:], "%d/%d", &pos1, &pos2)
			programs[pos1], programs[pos2] = programs[pos2], programs[pos1]
		case 'p':
			p1, p2 := cmd[1], cmd[3]
			pos1 := bytes.Index(programs, []byte{p1})
			pos2 := bytes.Index(programs, []byte{p2})
			programs[pos1], programs[pos2] = programs[pos2], programs[pos1]
		default:
			panic("unknown command: " + cmd)
		}
	}
	return string(programs)
}
