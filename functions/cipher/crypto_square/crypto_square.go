package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func ColsRows(n int) (c, r int) {
	c = int(math.Ceil(math.Sqrt(float64(n))))
	r = int(math.Ceil(float64(n) / float64(c)))
	return c, r
}

func Encode(s string) string {
	st := []string{}
	s = strings.ReplaceAll(strings.ToLower(s), " ", "")
	s = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
	c, r := ColsRows(len(s))

	for i:= 0; i < c ; i++{
		current := ""
		for j := i; j <= len(s)-1; j+=c{
			current+=string(s[j])
		}
		if len(current)-1 !=r{
			current += strings.Repeat(" ", r-len(current))
		}
		st = append(st, current)
	}
	return strings.Join(st, " ")
}

func main() {
	input := "See you later alligator!"
	fmt.Println(Encode(input))	
}
