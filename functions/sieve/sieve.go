package main

import (
	"fmt"
	"math"
)

func Find(limit int) (primes []int) {
	for current := 2; current <= limit; {
		if num, ok := Prime(current); ok {
			primes = append(primes, num)
		}
		current++
	}
	return primes
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
	fmt.Println(Prime(10))
	fmt.Println(Find(10))
}
