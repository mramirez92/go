package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Abbreviate(s string) (abv string) {
    s, abv = strings.ToUpper(s), string(s[0])
	for i:= 1; i < len(s); i++{
        if string(s[i]) != "'" && !unicode.IsLetter(rune(s[i])){
            if unicode.IsLetter(rune(s[i+1])){
                abv += string(s[i+1])
            }
        }
    }
    return 
}


func main(){
	s := "Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me"
	fmt.Println(Abbreviate(s))
}