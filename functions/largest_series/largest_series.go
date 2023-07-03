package main

import (
	"fmt"
	"strconv"
	// "errors"
)

func Series(code string, digits int) string {
	series := []string{}
	for i := 0; i <= len(code)-digits; i++ {
		series = append(series, code[i:i+digits])
	}
	return Largest(series)
}

func Largest(nums []string) string {
	current, lgProduct := "", 0

	for _, num := range nums {
		product := 1
		for _, char := range num {
			num, _ := strconv.Atoi(string(char))
			product *= num
		}
		if product > lgProduct {
			current, lgProduct = num, product
		}
	}
	return current
}

func main() {
	n := []string{"223", "999", "222"}
	nums := "63915"
	fmt.Println(Largest(n))

	fmt.Println(Series(nums, 3))
	// largest := ""
	// product := 1

	// for _, char := range n{
	// 	num, _ := strconv.Atoi(string(char))
	// 	// fmt.Println(num)
	// 	product *= num
	// }
	// fmt.Println(product)
}
