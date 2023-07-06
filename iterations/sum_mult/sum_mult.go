package main

import (
	"fmt"

)


func Multiples(level, baseVal int) []int{
	nums := []int{}
	for i := 1; ; i++ {
		if baseVal*i < level {
			nums = append(nums, baseVal*i)	
			}else{
				break
			}
	}
	return nums
}


func main() {
	fmt.Println(Multiples(20, 3))
}
