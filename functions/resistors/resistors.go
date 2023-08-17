package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var res = map[string]int{"black": 0, "brown": 1, "red": 2, "orange": 3, "yellow": 4, "green": 5, "blue": 6, "violet": 7, "grey": 8, "white": 9}

func Label(colors []string) string {
	num := ""
	for i := 0; i < 3; i++ {
		if i == 2 {
			num += strings.Repeat("0", res[colors[2]])

		} else {
			num += strconv.Itoa(res[colors[i]])
		}
	}
	n, _ := strconv.Atoi(num)
	return Ohms(n)
}

func Ohms(n int) string {
	abs := math.Abs(float64(n))

	if abs >= 1000 && abs < 1000000 {
		n = n / 1000
		return strconv.Itoa(n) + " kiloohms"
	} else if abs >= 1000000 && abs < 1000000000 {
		n = n / 1000000
		return strconv.Itoa(n) + " megaohms"
	} else if abs >= 1000000000 {
		n = n / 1000000000
		return strconv.Itoa(n) + " gigaohms"
	}
	return strconv.Itoa(int(abs)) + " ohms"
}

func main() {
	color := []string{"black", "yellow", "black"}
	f := Label(color)
	fmt.Println(f)

}
