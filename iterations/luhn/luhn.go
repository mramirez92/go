package luhn

import (
    "strconv"
    "strings"
    )

//  RevStrToInt takes a string, appends numerical characters as int to an array
func RevStrToInt(id string) []int {
    nums := make([]int, len(id))
    for i := range id {
        nums[i], _ = strconv.Atoi(string(id[len(id)-1-i]))
    }
	return nums
}

// ContainsLetters checks if a string contains alpha characters
func ContainsLetters(str string) bool {
	return strings.ContainsAny(str, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

// Valid confirms if a string is valid based on the Luhn formula 
func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
    if ContainsLetters(id) || len(id) == 1{
        return false
    }
	nums := RevStrToInt(id)
    sum := 0
    
    for i:=0; i < len(nums); i++{
        if i%2!=0{
            nums[i]*=2
            if nums[i] > 9{
                nums[i]-=9
            }
        }
    	sum += nums[i]
    }
	return sum %10 == 0
}