package main

import (
	"fmt"
	"unicode"
)

func Cipher(str string, key int) (code string) {
	shift := key % 26
	for _, char := range str {
		if unicode.IsLetter(char) {
			new := char + rune(shift)
			if unicode.IsLower(char) && new > 122 || unicode.IsUpper(char) && new > 90 {
				new -= 26
			} 
			code += string(new)
		} else {
			code += string(char)
		}
	}
	return code
}

func main() {
	s := "abcz ZX-12!ab3"
	fmt.Println(Cipher(s, 26))

}
