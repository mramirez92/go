package main

import (
	"errors"
	"fmt"
	"strings"
)

func Diamond(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("out of range byte")
	}
	if char == 'A' {
		return string(rune('A')), nil
	}

	width := (int(char)-64)*2 - 1

	l, r, c := 0, 0, (width - 2)

	letters := []string{}

	for i := char; i >= 'A'; i-- {
		letter := string(rune(i))
		if i == 'A' {
			row := strings.Repeat(" ", (width-1)/2) + letter + strings.Repeat(" ", (width-1)/2)
			letters = append([]string{row}, append(letters, row)...)
		} else {
			row := strings.Repeat(" ", l) + letter + strings.Repeat(" ", c) + letter + strings.Repeat(" ", r)
			letters = append([]string{row}, letters...)
			if i != char {
				letters = append(letters, row)
			}
			l, r, c = l+1, r+1, c-2
		}
	}
	return strings.Join(letters, "\n"), nil
}

func main() {
	fmt.Println(Diamond('M'))

}
