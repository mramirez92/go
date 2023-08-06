package main

import (
	"fmt"

	"strings"
)

func IsSubsequence(s, t string) bool {
	i := 0
	for _, char := range t {
		if i < len(s) && byte(char) == s[i] {
			i++
		}
	}
	return i == len(s)
}

func FindPivot(nums []int) int {
	left, right := 0, 0

	for _, num := range nums {
		right += num
	}

	for i, num := range nums {
		right -= num
		if right == left {
			return i
		}
		left += num
	}
	return -1
}

func MergeStrAlt(word1, word2 string) (merged string) {
	remainder, lenA, lenB := "", len(word1), len(word2)
	if lenA > lenB {
		remainder = word1[lenB:]
	} else {
		remainder = word2[lenA:]
	}

	for i, char := range word1 {
		if i < lenB {
			merged += string(char) + string(word2[i])
		}
	}
	return merged + remainder
}

func removeDuplicates(nums []int) int {
	seen := map[int]struct{}{}
	// unique := []int{}
	for _, num := range nums {
		if _, ok := seen[num]; !ok {
			seen[num] = struct{}{}
		}
	}
	return len(seen)
}

func remD(nums []int) int {
	seen := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != seen {
			nums[count] = nums[i]
			count++
		}
		seen = nums[i]
	}
	return count
}

func flower(spaces []int, new int) bool {
	count := 0
	for i := 0; i < len(spaces); i++ {
		if spaces[i] == 0 && (i == 0 || spaces[i-1] == 0) && (i == len(spaces)-1 || spaces[i+1] == 0) {
			count++
			spaces[i] = 1
		}
	}
	return count >= new
}

func MaxSub(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, current := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if current+nums[i] > nums[i] {
			current += nums[i]
		} else {
			current = nums[i]
		}
		if current > max {
			max = current
		}
	}
	return max
}

func maxPro(prices []int) (profit int) {
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return
}

func sorted(n []int) (max int) {
	max = n[0]
	for _, num := range n {
		if num > max {
			max = num
		}
	}
	return
}

func candy(nums []int, add int) []bool {
	max := sorted(nums)

	greater := []bool{}

	for _, num := range nums {
		if num+add >= max {
			greater = append(greater, true)
		} else {
			greater = append(greater, false)
		}
	}
	return greater
}

// *helper function for reversing vowel in string
func isVowel(char byte) bool {
	char = strings.ToLower(string(char))[0]
	return char == 'a' || char == 'e' || char == 'o' || char == 'i' || char == 'u'
}

func RevVowels(s string) string {
	str := []byte(s)
	i, j := 0, len(s)-1

	for i < j {
		if !isVowel(str[i]) {
			i++
		} else if !isVowel(str[j]) {
			j--
		} else {
			str[i], str[j] = str[j], str[i]
			i++
			j++
		}
	}
	return string(str)
}


func main() {
	fmt.Println("hello world")

}
