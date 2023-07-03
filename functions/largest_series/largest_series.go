package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Series(digits string, span int) (current int64, e error) {
	if span < 0 {
		return 0, errors.New("span must not be negative")
	} else if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}
	nums := strings.Split(digits, "")

	for i := 0; i <= len(nums)-span; i++ {
		var product int64 = 1
		for j := 0; j < span; j++ {
			if num, e := strconv.Atoi(nums[i+j]); e != nil {
				return 0, errors.New("digits input must only contain digits")
			} else {
				product *= int64(num)
			}
		}
		if product > current {
			current = product
		}
	}
	return current, e
}

func main() {
	nums := "29"
	fmt.Println(Series(nums, 2))

}
