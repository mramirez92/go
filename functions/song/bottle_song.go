package main

import (
	"fmt"
	"strings"
)

var song = "%s green bottles hanging on the wall,*%s green bottles hanging on the wall,*And if one green bottle should accidentally fall,*There'll be %s green bottles hanging on the wall.*"
var nums = map[int]string{0: "no", 1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 10: "Ten"}

func Recite(start, lines int) (l []string) {
	for i := 1; i <= lines; i++ {
		lyrics := song
		if lines == 1 || i == lines {
			lyrics = song[:(len(song) - 1)]
		}
		if start == 2 {
			lyrics = strings.ReplaceAll(lyrics, "bottles hanging on the wall.", "bottle hanging on the wall.")
		}
		if start == 1 {
			lyrics = strings.Replace(strings.Replace(lyrics, "bottles", "bottle", 3), "bottle hanging on the wall.", "bottles hanging on the wall.", -1)
		}
		mod := strings.Split(fmt.Sprintf(lyrics, nums[start], nums[start], strings.ToLower(nums[start-1])), "*")
		l = append(l, mod...)
		start--
	}
	return l
}

func main() {
	lines := Recite(10, 2)
	fmt.Println(lines)
}
