package main

import "fmt"

type Link struct {
	l, r int
}

type Linked struct {
	Link
	r int
}

func day24r(links []Link, hist []Linked, f func(hist []Linked)) {
	to := 0
	if len(hist) > 0 {
		to = hist[len(hist)-1].r
	}
	var found bool
	for i, link := range links {
		var r int
		switch {
		case link.l == to:
			r = link.r
		case link.r == to:
			r = link.l
		default:
			continue
		}
		found = true
		dlinks := CopySlice(links)
		dlinks[i], dlinks[len(dlinks)-1] = dlinks[len(dlinks)-1], dlinks[i]
		dlinks = dlinks[:len(dlinks)-1]
		day24r(dlinks, append(hist, Linked{link, r}), f)
	}
	if !found {
		f(hist)
	}
}

func day24score(hist []Linked) (sum int) {
	for _, h := range hist {
		sum += h.Link.l + h.Link.r
	}
	return
}

func day24parse(input []string) []Link {
	links := make([]Link, len(input))
	for i, line := range input {
		fmt.Sscanf(line, "%d/%d", &links[i].l, &links[i].r)
	}
	return links
}

func day24a(input []string) int {
	var maxscore int
	day24r(day24parse(input), nil, func(hist []Linked) {
		score := day24score(hist)
		if score > maxscore {
			maxscore = score
		}
	})
	return maxscore
}

func day24b(input []string) int {
	var maxlen, maxscore int
	day24r(day24parse(input), nil, func(hist []Linked) {
		if len(hist) > maxlen {
			maxscore = day24score(hist)
			maxlen = len(hist)
		}
		if len(hist) == maxlen {
			score := day24score(hist)
			if score > maxscore {
				maxscore = score
			}
		}
	})
	return maxscore
}
