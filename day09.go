package main

func day9a(input string) (total int) {
	total, _ = day9(input)
	return
}

func day9b(input string) (gb int) {
	_, gb = day9(input)
	return
}

func day9(input string) (total int, gb int) {
	depth := 1
	var garbage bool
	for i := 0; i < len(input); i++ {
		chr := input[i]

		if garbage {
			switch chr {
			case '!':
				i++
			case '>':
				garbage = false
			default:
				gb++
			}
			continue
		}

		switch chr {
		case '<':
			garbage = true
		case '{':
			total += depth
			depth++
		case '}':
			depth--
		}
	}
	return
}
