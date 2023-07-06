package main

import (
	"fmt"
)

type IntList []int

// **foldl

//**foldr

// ?
func (s IntList) Filter(fn func(int) bool) (match IntList) {
	odd := func(n int) bool {
		return n%2 == 1
	}
	for _, num := range s {
		if odd(num) {
			match = append(match, num)
		}
	}
	return
}

func (s IntList) Length() int {
	return len(s)
}

// ** map

// ** reverse**

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = append(s, list...)
	}
	return s
}

// type.method(changes?)

func main() {
	s := []IntList{[]int{3}, []int{}, []int{4, 5, 6}}
	n := IntList{1}

	// fmt.Println(s.Append(n))
	// fmt.Println(s.Filter(func(i int) bool {
	// 	return i %2 ==1
	// }))

	fmt.Println(n.Concat(s))

}
