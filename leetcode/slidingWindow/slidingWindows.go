package main

import (
	"fmt"
	"math"
)

func findMaxAvg(nums []int, k int) float64 {
	max := math.Inf(-1)
	for i := 0; i <= len(nums)-k; i++ {
		current := 0.0
		for j := i; j < i+k; j++ {
			current += float64(nums[j])
		}
		current /= float64(k)
		if current > max {
			max = current
		}
	}
	return max
}

func maxVowels(s string, k int) int {
	max, current := 0, 0
	vowels := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}

	for i := 0; i < len(s); i++ {
		if _, ok := vowels[s[i]]; ok {
			current++
		}

		if i >= k-1 {
			if current > max {
				max = current
			}

			if _, ok := vowels[s[i-k+1]]; ok {
				current--
			}
		}
	}
	return max
}

func main() {
	// nums, k := []int{1,12,-5,-6,50,3},  4
	fmt.Println(findMaxAvg([]int{-1}, 1))

	s, k := "weallloveyou", 7
	fmt.Println(maxVowels(s, k))
}
