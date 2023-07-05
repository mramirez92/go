package main

import "fmt"

func factors(n int) (primes []int) {
	for i := 2; i <= n; {
		if n%i == 0 {
			primes = append(primes, i)
			n /= i
		} else {
			i++
		}
	}
	return primes
}

func main() {
	num := 9
	fmt.Println(factors(num))
}
