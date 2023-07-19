package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type shift int
type vigenere string
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

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

func Atbash(s string) string {
	coded := []string{}
	new := ""
	for _, char := range strings.ToLower(s) {

		if len(new) < 5 {
			if unicode.IsLetter(char) {
				r := 122 - (int(char) - 97)
				new += string(rune(r))
			} else if unicode.IsNumber(char) {
				new += string(char)
			}
		}
		if len(new) == 5 {
			coded = append(coded, new)
			new = ""
		}
	}
	if len(new) > 0 {
		coded = append(coded, new)
	}
	return strings.Join(coded, " ")
}

func (c shift) Encode(input string) string {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
	new := ""
	for _, char := range input {
		if unicode.IsLetter(char) {
			char += rune(c)
			if char > 122 {
				char = (char - 122) + 96
				new += string(char)
			} else if char < 97 {
				char = 123 - (97 - char)
				new += string(char)
			} else {
				new += string(char)
			}
		}
	}
	return new
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}
	return shift(distance)
}
func (c shift) Encode(input string) (new string) {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
	for _, char := range input {
		if unicode.IsLetter(char) {
			char += rune(c)
			if char > 122 {
				char = (char - 122) + 96
				new += string(char)
			} else if char < 97 {
				char = 123 - (97 - char)
				new += string(char)
			} else {
				new += string(char)
			}
		}
	}
	return new
}

func (c shift) Decode(input string) string {
	decodeShift := shift(-int(c))
	return decodeShift.Encode(input)
}

func NewVigenere(key string) Cipher {
	alphaOnly := regexp.MustCompile("[^a-z]").ReplaceAllString(key, "")
	if key != strings.ToLower(key) || strings.Count(key, "a") == len(key) || len(alphaOnly) != len(key) {
		return nil
	}
	return vigenere(key)
}

func (v vigenere) Encode(input string) (new string) {
	result := regexp.MustCompile("[^a-z]").ReplaceAllString(strings.ToLower(input), "")
	keyLen := len(v)
	for i, char := range result {
		if unicode.IsLetter(char) {
			val := shift(v[i%keyLen] - 'a')
			new += shift(val).Encode(string(char))
		}
	}
	return new
}

func (v vigenere) Decode(input string) (new string) {
	key := strings.ToLower(strings.ReplaceAll(string(v), " ", ""))
	keyLen := len(key)
	for i, char := range input {
		if unicode.IsLetter(char) {
			val := shift(key[i%keyLen] - 'a')
			deshift := shift(-int(val))
			new += deshift.Encode(string(char))
		}
	}
	return new
}

func main() {
	// s := "ab1 zyx 5678,"
	// fmt.Println(Cipher(s, 26))
	// fmt.Println(Atbash(s))
	// t := shift(2)
	// fmt.Println(t. Encode(s))

	key := vigenere("qgbvno")
	fmt.Println(key.Encode("cof-FEE,_123!"))

}
