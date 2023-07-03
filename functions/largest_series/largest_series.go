package main

import (
	"fmt"
	"strconv"
	"errors"
)

func Series(digits string, span int) (int64, error) {
	switch{
	case span <0:
		return 0, errors.New("span must not be negative")
	case span >len(digits):
		return 0, errors.New("span must be smaller than string length")
	}
	series := []string{}
	for i := 0; i <= len(digits)-span; i++ {
		series = append(series, digits[i:i+span])
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
	nums := "12345"
	fmt.Println(Series(nums, -1))
	
}
