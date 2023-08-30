package main

import "fmt"

func Counts(dice []int) (map[int]int, int) {
	rollCount, sum := map[int]int{}, 0
	for _, num := range dice {
		rollCount[num]++
		sum += num
	}
	return rollCount, sum
}

func Score(dice []int, category string) int {
	rollCount, sum := Counts(dice)
	numCat := map[string]int{"ones": 1, "twos": 2, "threes": 3, "fours": 4, "fives": 5, "sixes": 6}
	v, ok := numCat[category]; if ok{
		count, ok := rollCount[v]; if ok{
			return v * count
		}
		// if ok {
		// 	return v * count
		// }	
	}
	// if ok {
	// 	count, ok := rollCount[v]
	// 	if ok {
	// 		return v * count
	// 	}
	// }
	count := len(rollCount)

	if category==  "choice"{
		return sum
	}else if category == "yacht" && count == 1 {
		return 50
	}else if category == "four of a kind"{
		for k, v := range rollCount {
			if v ==  1 && count == 2 || v == 5{
				return sum - k
	}
}
	}else if category == "full house"{
    		for _, v := range rollCount {
			if v == 3 && count == 2{
				return sum 
	}
            }
    
	}else if count == 5 && category == "big straight" && sum == 20 {
		return 30
	} else if count == 5 && category == "little straight" && sum == 15 {
		return 30
	}
	return 0
}

func main() {
	r := []int{4, 4, 4, 4, 4}
	// ls := []int{2, 2, 2, 2, 2}
	fmt.Println(Score(r, "four of a kind"))
	fmt.Println(Counts(r))
	// fmt.Println(Score(ls, "yacht"))
}
