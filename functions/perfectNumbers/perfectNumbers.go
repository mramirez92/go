package main

import (
	"errors"
	"fmt"
)

type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive error = errors.New("input is less than 0")

func Classify(num int64) (Classification, error) {
	if num <= 0 {
		return 0, ErrOnlyPositive
	}

	sum := AliquotSum(num)
	switch {
	case sum < num:
		return ClassificationDeficient, nil
	case sum == num:
		return ClassificationPerfect, nil
	default:
		return ClassificationAbundant, nil
	}
}
func AliquotSum(num int64) int64 {
	sum := int64(0)
	for i := int64(1); i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(Classify(1))
	fmt.Println(AliquotSum(1))
}
