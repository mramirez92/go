package main

import "fmt"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := IntList{}
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	res := IntList{}
	for _, v := range s {
		res = append(res, fn(v))
	}
	return res
}

func (s IntList) Reverse() IntList {
	res := IntList{}
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return res
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = append(s, list...)
	}
	return s
}

func main() {
	// s := []IntList{[]int{3}, []int{}, []int{4, 5, 6}}
	s := IntList{1, 2}

	// fmt.Println(s.Append(n))
	fmt.Println(s.Filter(func(i int) bool {
		return i%2 == 0
	}))

	fmt.Println(s.Map(func(i int) int {
		return i + 1
	}))

}
