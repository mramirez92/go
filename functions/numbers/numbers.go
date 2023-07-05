package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func Armstrong(n int) bool {
	num, total := strconv.Itoa(n), 0
	for _, char := range num {
		digit, _ := strconv.Atoi(string(char))
		total += int(math.Pow(float64(digit), float64(len(num))))
	}
	return total == n
}

func Number(n string) (num string, e bool) {
	for _, char := range n {
		if unicode.IsDigit(char) {
			num += string(char)
		}
	}
	l := len(num)
	if l == 11 && num[0] == '1' {
		num, l = num[1:], 10
	}
	if l == 10 {
		if num[0] == '1' || num[0] == '0' || num[3] == '1' || num[3] == '0' {
			return "", true
		} else {
			return num, false
		}
	}
	return "", true
}

func AreaCode(phone string) (string, bool) {
	if n, er := Number(phone); er {
		return "", true
	} else {
		return n[:3], er
	}
}

func Format(phone string) (string, bool) {
	if n, er := Number(phone); er {
		return "", true
	} else {
		return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), false
	}

}

func main() {
	num := "223 45a 1234"
	fmt.Println(Number(num))
	fmt.Println(AreaCode(num))
	fmt.Println(Format(num))
}
