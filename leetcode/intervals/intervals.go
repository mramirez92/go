package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	start, end := "", ""
	ranges := []string{}
	for i := 0; i < len(nums); {
		start = strconv.Itoa(nums[i])
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			ranges = append(ranges, start)
			i++
		} else {
			for j := i + 1; j < len(nums); j++ {
				if nums[j-1]+1 != nums[j] {
					end = strconv.Itoa(nums[j-1])
					ranges = append(ranges, fmt.Sprintf("%s->%s", start, end))
					i = j
					break
				}
			}
		}
	}
	return ranges
}


func sumR(nums []int) []string {
	ranges := []string{}
	for i := 0; i < len(nums); {
		start := strconv.Itoa(nums[i])
		j := i + 1
		for ; j < len(nums) && nums[j-1]+1 == nums[j]; j++ {
		}

		if j-1 > i {
			ranges = append(ranges, fmt.Sprintf("%s->%s", start, strconv.Itoa(nums[j-1])))
		} else {
			ranges = append(ranges, start)
		}

		i = j
	}

	return ranges
}

func main() {
	nums := []int{0,1,2,4,5,7}
	fmt.Println(sumR(nums))
}
