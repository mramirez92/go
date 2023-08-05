package main

import (
	"fmt"
)

func Flat(nested []interface{}) []interface{}{
	new := []interface{}{}
	for _, item := range nested{
		if num, ok := item.(int); ok{
			new = append(new, num)
		}else{
			new = append(new, CheckNested(item)...)
		}
	}
	return new
}

func CheckNested(item interface{})[]interface{}{
	nums := []interface{}{}
	switch val := item.(type){
	case int:
		nums = append(nums, val)
	case []interface{}:
		for _, nestedEle := range  val{
			nums = append(nums, CheckNested(nestedEle)...)
		}
	}
	return nums
}


func main(){
	// arr := []interface{}{1, []interface{}{2, 3, 4, 5, 6, 7}, 8}
	levelFive := []interface{}{0, 2, []interface{}{[]interface{}{2, 3}, 8, 100, 4, []interface{}{[]interface{}{[]interface{}{50}}}}, -2}

	fmt.Println(Flat(levelFive))
}