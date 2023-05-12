package leetcode

func IsSubsequence(s, t string) bool{
	i := 0
	for _, char  := range t{
		if i < len(s) &&  byte(char) == s[i]{
			i++	
		}
	}
	return i == len(s)
}

func FindPivot(nums [] int) int{
	left, right := 0 , 0

	for _, num := range nums { 
		right += num 
	}

	for i, num := range nums {
		right -= num 
		if right == left{
			return i
		}
		left += num 
	}
	return -1
}