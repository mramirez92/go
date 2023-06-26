package main

import "fmt"

var alpha = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
  }

func Cipher(str string, key int) string {
	new := ""
	for _, char := range str {
		found := false
		for i, r := range alpha {
			if char == r {
				i += key
				if i > 4 {
					i -= 5
				}
				new += string(alpha[i])
				found = true
			}
		}
		if !found {
			new += string(char)
		}
	}
	return new
}

func main(){
	s := "abcd E"
	fmt.Println(Cipher(s,2))

}