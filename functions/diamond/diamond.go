package main

import (
	"fmt"
	"strings"
	"errors"
)

func Diamond(char byte) (string, error) {
    if char <'A' || char > 'Z'{
        return "", errors.New("out of range byte")
    }
	if char == 'A' {
		return string(rune('A')), nil
    }
	middle := int(char)

	n := (middle-64)*2 - 1

	l, r, c := 0, 0, (n - 2)

	letters := []string{}

	for i := middle; i >= 65; i-- {
		if i == 65 {
			row := strings.Repeat(" ", (n-1)/2) + string(rune(i)) + strings.Repeat(" ", (n-1)/2)
			letters = append([]string{row}, letters...)
		} else {
			row := strings.Repeat(" ", l) + string(rune(i)) + strings.Repeat(" ", c) + string(rune(i)) + strings.Repeat(" ", r)
			letters = append([]string{row}, letters...)
			l++
			r++
			c -= 2
		}
	}

	for i := len(letters) - 2; i >= 0; i-- {
		letters = append(letters, letters[i])
	}
	return strings.Join(letters, "\n"), nil
}

func main(){
	fmt.Println(Diamond('c'))

	
}