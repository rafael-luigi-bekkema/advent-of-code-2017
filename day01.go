package main

func day1a(digits string) int {
	var total int
	for i := range []byte(digits) {
		nexti := (i + 1) % len(digits)
		if digits[i] == digits[nexti] {
			total += int(digits[i] - '0')
		}
	}
	return total
}

func day1b(digits string) int {
	var total int
	for i := range []byte(digits) {
		nexti := (i + (len(digits) / 2)) % len(digits)
		if digits[i] == digits[nexti] {
			total += int(digits[i] - '0')
		}
	}
	return total
}
