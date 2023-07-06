package main

import (
	"fmt"
)

type Relation string

// Possible relations
const (
	RelationEqual     Relation = "equal"
	RelationSublist   Relation = "sublist"
	RelationSuperlist Relation = "superlist"
	RelationUnequal   Relation = "unequal"
)

func Sublist(a, b []int) Relation {
	la, lb := len(a), len(b)
	switch {
	case la == lb && Contains(a, b):
		return RelationEqual
	case la < lb && Contains(a, b):
		return RelationSublist
	case la > lb && Contains(b, a):
		return RelationSuperlist
	}
	return RelationUnequal
}

func Contains(short, long []int) bool {
	for i := 0; i <= len(long)-len(short); i++ {
		match := true
		for j := 0; j < len(short); j++ {
			if short[j] != long[i+j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func main() {

	b := []int{1, 2, 3}
	a := []int{1, 2, 3}

	fmt.Println(Sublist(a, b))
}
