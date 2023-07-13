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

var log = map[rune]string{'❗': "recommendation", '🔍': "search", '☀': "weather"}

// Application identifies the application emitting the given log.
func Application(s string) string {
    for _, char := range s {
        if val, ok := log[char]; ok{
            return val
        }
    }
    return "default"
}


func main(){
	s := "❗ help"
	fmt.Println(Application(s))
}