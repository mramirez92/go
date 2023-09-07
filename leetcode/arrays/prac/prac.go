package main

import "fmt"

func zeros(nums []int) []int {
	for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            for j := i + 1; j < len(nums); j++ {
                if nums[j] != 0 {
                    nums[i], nums[j] = nums[j], nums[i]
                    break
                }
            }
        }
    }
    return nums
}


func main(){
	num := []int{2, 0,3,0,4,}
	fmt.Println(zeros(num))
}