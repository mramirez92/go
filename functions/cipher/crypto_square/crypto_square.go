package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func filter(s string) string {
	s = strings.ReplaceAll(strings.ToLower(s), " ", "")
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
}

func sides(n int) (c, r int) {
	square := int(math.Sqrt(float64(n)))
	c, r = square, square 
	if (c * r) >= n {
		return
	}
	return c + 1, r
}

func split(s string) []string {
	st := []string{}
	s = filter(s)
	c, _ := sides(len(s))

	for i:= 0; i < c ; i++{
		current := ""
		for j := i; j <= len(s)-1; j+=c{
			current+=string(s[j])
		}
		st = append(st, current)
	}
	return st
}

func main() {
	input := "123456789a"
	// l :=filter(input)
	// fmt.Println(len(l))
	// fmt.Println(sides(len(l)))
	fmt.Println(split(input))

	
}
