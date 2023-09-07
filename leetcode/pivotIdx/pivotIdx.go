package main

import "fmt"

func pivIdx(nums []int)int{
	right, left := 0, 0
	for _, n := range nums{
		right+= n
	}

	for i, n := range nums{
		right -= n
		if left == right{
			return i
		}
		left += n
	}
	return -1
}

// 2215. Find the Difference of Two Arrays
func findDifference(nums1 []int, nums2 []int) [][]int {
	s1, s2 := makeMap(nums1), makeMap(nums2)
	diff1, diff2 := []int{}, []int{}

	for k := range s1{
		if !s2[k]{
			diff1 = append(diff1, k)
		}
	}
	for k := range s2{
		if !s1[k]{
			diff2 = append(diff2, k)
		}
	}
	return [][]int{diff1, diff2}
}

//  *** helper function to create map of nums for findDifference ***
func makeMap(nums []int)map[int]bool{
	n := map[int]bool{}
	for _, num := range nums{
		if _, ok := n[num]; !ok{
			n[num] = true
		}
	}
	return n
}

func main(){
	n := []int{1,7,3,6,5,6}
	n1, n2 := []int{1,2,3,3}, []int{4,6}
	fmt.Println(pivIdx(n))

	fmt.Println(makeMap(n2))
	fmt.Println(findDifference(n1, n2))

}