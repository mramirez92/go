package main

import (
	"fmt"
	"strconv"
	"strings"
)

var ops = map[string]func(int, int) int{
	"plus":       func(a, b int) int { return a + b },
	"minus":      func(a, b int) int { return a - b },
	"multiplied": func(a, b int) int { return a * b },
	"divided":    func(a, b int) int { return a / b },
}

func Answer(q string) (int, bool) {
	words := []string{}
	if strings.HasPrefix(q, "What is") && strings.HasSuffix(q, "?") {
		q = strings.ReplaceAll(q, "by", "")
		words = strings.Fields(strings.TrimSuffix(strings.TrimPrefix(q, "What is"), "?"))
		if len(words) > 0 {
			return Math(words)
		}
	}
	return 0, false
}

func Math(w []string) (int, bool) {
	// First element isn't a number, return false
	total, e := strconv.Atoi(w[0])
	if e != nil {
		return 0, false
	}

	for i := 1; i < len(w); i += 2 {
		if op, ok := ops[w[i]]; ok {
			if i+1 >= len(w) {
				return 0, false
			}
			// Check next element is a number
			if num, e := strconv.Atoi(w[i+1]); e == nil {
				total = op(total, num)
			} else {
				return 0, false
			}
		} else {
			return 0, false
		}
	}

	return total, true
}

func main() {
	s := "What is 5?"
	result, success := Answer(s)
	fmt.Println(result, success) // Output: 4 true
}
