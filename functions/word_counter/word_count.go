package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`^\W*(\w+(?:\W+\w+)*).*\W*$`)
	words := map[string]int{}
	split := regexp.MustCompile(`[,\s\\]+`).Split(strings.ToLower(phrase), -1)
	for _, word := range split {
		if word != "" {
			w := re.ReplaceAllString(word, "$1")
			if _, ok := words[w]; ok {
				words[w] = words[w] + 1
			} else {
				words[w] = 1
			}
		}
	}
	return words
}