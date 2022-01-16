package main

import "fmt"

/*
Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.
*/

func day25a(input []string) int {
	var state byte
	fmt.Sscanf(input[0], "Begin in state %c.", &state)
	var steps int
	fmt.Sscanf(input[1], "Perform a diagnostic checksum after %d steps.", &steps)

	type Action struct {
		Value uint8
		MoveR bool
		State byte
	}
	states := map[byte][2]Action{}

	i := 2
	for i < len(input) {
		var fstate, tstate1, tstate2 byte
		var val1, val2 uint8
		var dir1, dir2 string
		fmt.Sscanf(input[i+1], "In state %c:", &fstate)
		fmt.Sscanf(input[i+3], " - Write the value %d.", &val1)
		fmt.Sscanf(input[i+4], " - Move one slot to the %s.", &dir1)
		fmt.Sscanf(input[i+5], " - Continue with state %c.", &tstate1)

		fmt.Sscanf(input[i+7], " - Write the value %d.", &val2)
		fmt.Sscanf(input[i+8], " - Move one slot to the %s.", &dir2)
		fmt.Sscanf(input[i+9], " - Continue with state %c.", &tstate2)

		states[fstate] = [2]Action{
			{val1, dir1 == "right.", tstate1},
			{val2, dir2 == "right.", tstate2},
		}

		i += 10
	}

	var cursor int
	tape := map[int]uint8{}
	for i := 0; i < steps; i++ {
		action := states[state][tape[cursor]]
		tape[cursor] = action.Value
		if action.MoveR {
			cursor++
		} else {
			cursor--
		}
		state = action.State
	}
	var sum int
	for _, t := range tape {
		if t == 1 {
			sum++
		}
	}
	return sum
}
