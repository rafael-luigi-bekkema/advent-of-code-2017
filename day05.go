package main

func day5a(offsets []int) (steps int) {
	offsets = CopySlice(offsets)
	var ptr int
	for 0 <= ptr && ptr < len(offsets) {
		offsets[ptr], ptr = offsets[ptr]+1, ptr+offsets[ptr]
		steps++
	}
	return steps
}

func day5b(offsets []int) (steps int) {
	offsets = CopySlice(offsets)
	var ptr int
	for 0 <= ptr && ptr < len(offsets) {
		offset := offsets[ptr]
		if offset >= 3 {
			offsets[ptr]--
		} else {
			offsets[ptr]++
		}
		ptr += offset
		steps++
	}
	return steps
}
