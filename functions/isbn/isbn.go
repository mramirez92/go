package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Valid(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	re := regexp.MustCompile(`^([0-9Xx]{10})$`)
	total := 0

	if re.MatchString(isbn) {
		for i, char := range isbn {
			if i == 9 && (char == 'X' || char == 'x') {
				total += 10 * (10 - i)
			} else {
				total += int(char - '0') * (10 - i)
			}
		}
		return total%11 == 0
	}
	return false
}

func main() {
	i := "359821507X"
	fmt.Println(Valid(i))

}