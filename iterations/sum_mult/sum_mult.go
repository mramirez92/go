package main

import (
	"fmt"
)

func SumMultiples(limit int, divisors ...int) (total int) {
	mult := map[int]struct{}{}
	for _, val := range divisors {
		if val != 0 {
			current := val
			for current < limit {
				if _, ok := mult[current]; !ok {
					mult[current] = struct{}{}
					total += current
				}
				current += val
			}
		}
	}
	return
}

func main() {
	fmt.Println(SumMultiples(7, 3))
}
