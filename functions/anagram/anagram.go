package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) (ana []string) {
	for _, word := range candidates {
		if strings.ToLower(subject) != strings.ToLower(word) {
			subject = Sort(subject)
			if subject == Sort(word) {
				ana = append(ana, word)
			}
		}
	}
	return ana
}

func Sort(word string) string {
	chars := strings.Split(strings.ToLower(word), "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}