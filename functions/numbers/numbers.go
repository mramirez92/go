package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
	// "reflect"
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
	switch {
	case l > 11 || l < 10:
		return "", true
	case l == 11:
		if num[0] == '1' {
			if num[1] == '1' || num[1] == '0' || num[4] == '1' || num[4] == '0'{
				return "", true
			}
			num = num[1:]
		}else{
			return "", true
		}
	case l == 10:
		if num[0] == '1' || num[0] == '0' || num[3] == '1' || num[3] == '0'{
			return "", true
		}
	}
	return num, e

}

func AreaCode(phone string) (string, bool){
	if n, er := Number(phone); er{
		return "", true
	}else{
		return n[:3], er
	}
}

func Format(phone string)(string, bool){
	if n, er := Number(phone); er{
		return "", true
	}else{
		return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), false
	}

}



func main() {
	num := "223 256 9231"
	fmt.Println(Number(num))
	fmt.Println(AreaCode(num))
	fmt.Println(Format(num))

	// if num[1] == '1'{
	// 	fmt.Println("match")
	// }else{
	// 	fmt.Println("no match")
	// }
}
