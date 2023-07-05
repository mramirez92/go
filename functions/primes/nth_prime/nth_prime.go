package main

import (
	"errors"
	"fmt"
	"math"
)

func Find(nth int) (res int, e error) {
	if nth <= 0 {
		return 0, errors.New("input can't be less than 0")
	}
	current := 2
	for i := 1; i <= nth; {
		if num, ok := Prime(current); ok {
			res = num
			i++
		}
		current++
	}
	return res, e
}

func Prime(n int) (int, bool) {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return n, false
		}
	}
	return n, true
}

func main() {
	fmt.Println(Find(0))

}
