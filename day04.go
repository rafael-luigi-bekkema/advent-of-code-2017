package main

import (
	"sort"
	"strings"
)

func day4a(words []string) (result int) {
outer:
	for _, word := range words {
		swords := strings.Split(word, " ")
		pwords := make([]string, 0, len(swords))
		for _, sword := range swords {
			if In(sword, pwords...) {
				continue outer
			}
			pwords = append(pwords, sword)
		}
		result++
	}
	return
}

func day4b(words []string) (result int) {
outer:
	for _, word := range words {
		swords := strings.Split(word, " ")
		pwords := make([]string, 0, len(swords))
		for _, sword := range swords {
			ssword := []byte(sword)
			sort.Slice(ssword, func(i, j int) bool {
				return ssword[i] < ssword[j]
			})
			sssword := string(ssword)
			if In(sssword, pwords...) {
				continue outer
			}
			pwords = append(pwords, sssword)
		}
		result++
	}
	return
}
