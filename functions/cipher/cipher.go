package main

import (
	"fmt"
	"unicode"
)

func Ceasar(str string, key int) (code string) {
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

func Atbash(str string) []string {
	coded := []string{}
	new := ""
	for _, char := range str {
			if len(new) < 5{
				if unicode.IsLetter(char) {
					r := 122 - (int(char) - 97)
					new += string(rune(r))
				} else if unicode.IsNumber(char) {
					new += string(char)
				}
			}
			if len(new) == 5{
				coded = append(coded, new)
				new =""
			}
			}
			if len(new) > 0 {
				coded = append(coded, new)
			}
			return coded
		}



func main() {
	s := "ab1 zyx 5678,"
	// fmt.Println(Cipher(s, 26))

	fmt.Println(Atbash(s))

}
