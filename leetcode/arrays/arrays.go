package main

import (
	"fmt"
)

func Max(candies []int, extra int) (greatest []bool) {
	max := candies[0]
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}
	for _, num := range candies {
		if num+extra >= max {
			greatest = append(greatest, true)
		} else {
			greatest = append(greatest, false)
		}
	}
	return greatest
}

func MaxOperations(nums []int, k int) int {
	remainders := map[int]int{}
	pairs := 0
	for _, num := range nums {
		if remainders[num] == 0 {
			remainders[k-num]++
		} else {
			remainders[num]--
			pairs++
		}
	}
	return pairs
}

func LargestAltitude(gain []int) int {
	max, current := 0, 0
	for _, num := range gain {
		current += num
		if current > max {
			max = current
		}
	}
	return max
}

// ** 238. Product of Array Except Self **
// * Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
func productExceptSelf(nums []int) []int {
	products := []int{}
	previous := 1

	for i := 0; i < len(nums); i++ {
		res := 1
		for j := i + 1; j < len(nums); j++ {
			res *= nums[j]
		}
		products = append(products, (res * previous))
		previous *= nums[i]
		res = 1
	}
	return products
}

func main() {
	l := []int{1, 2, 3,4}

	// fmt.Println(Max(l, 1))
	// fmt.Println(MaxOperations(l, 5))
	// gain := []int{-4, -3, -2, -1, 4, 3, 2}
	// fmt.Println(LargestAltitude(gain))
	fmt.Println(productExceptSelf(l))
}
