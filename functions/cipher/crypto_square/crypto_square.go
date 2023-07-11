package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func filter(s string) string {
	s = strings.ReplaceAll(strings.ToLower(s), " ", "")
	// reg := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	// filtered := reg.ReplaceAllString(s, "")
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

func main() {
	// input := "Hello! 123, "

	// fmt.Println(filter(input)) // Output: Hello123
	fmt.Println(sides(54))
}
