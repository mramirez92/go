package main

import (
	"errors"
	"fmt"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}else if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}
	exp := len(inputDigits)-1
	res := 0
	for _, num := range inputDigits{
        if num >= inputBase || num < 0{
            return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
        }
		res += num * int(math.Pow(float64(inputBase), float64(exp)))
		exp--
	}
	if len(inputDigits) == 0 || res == 0{
        return []int{0}, nil
    }
	return Output(res, outputBase)
}
func Output(num int, outputBase int)(output []int, e error){
	for num > 0 {
		res := num % outputBase 
			output = append([]int{res}, output... )
			num = num / outputBase
	}
	return 	
}

func main() {
	inputBase := 2
	inputDigits := []int{1, 0, 1, 0, 1, 0}
	outputBase := 10

	fmt.Println(ConvertToBase(inputBase, inputDigits, outputBase))
	fmt.Println(ConvertToBase(10, []int{0}, 2))

}
