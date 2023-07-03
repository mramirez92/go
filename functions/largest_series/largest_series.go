package main

import (
	"fmt"
	"strconv"
	"errors"
)

func Series(code string, digits int) (int64, error) {
	series := []string{}
	for i := 0; i <= len(code)-digits; i++ {
		series = append(series, code[i:i+digits])
	}
	return Largest(series)
}

func Largest(nums []string) (current int64, e error) {
	lgProduct :=  0

	for _, num := range nums {
		product := 1
		for _, char := range num {
			if num, e := strconv.Atoi(string(char)); e != nil{
				return 0, errors.New("digits input must only contain digits")
			}else{
				product *= num
			}
		}
		if product > lgProduct {
			n, _ := strconv.Atoi(num)
			current, lgProduct = int64(n), product
		}
	}
	return current, e
}

func main() {
	nums := "63915"
	fmt.Println(Series(nums, 3))
	
}
