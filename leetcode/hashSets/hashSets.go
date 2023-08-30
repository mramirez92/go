package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// *helper function for closeStrings, returns map of letters and slice of letter counts sorted

func count(word string) (map[string]struct{}, []int) {
    letters := map[string]struct{}{}
    counts := []int{}
    for _, char := range word {
        if _, ok := letters[string(char)]; !ok {
            letters[string(char)] = struct{}{}
            counts = append(counts, strings.Count(word, string(char)))
        }
    }
    sort.Ints(counts)
    return letters, counts
}

func closeStrings(word1, word2 string) bool {
    w1, wCount := count(word1)
    w2, w2Count := count(word2)

    return reflect.DeepEqual(w1, w2) && reflect.DeepEqual(wCount, w2Count)
}

// ransom note

func wordMap(word string)map[rune]int{
	runes := map[rune]int{}
	for _, char := range word{
		if _, ok := runes[char]; !ok{
			runes[char] = strings.Count(word, string(char))
		}
	}
	return runes
}

func RansomeNote(note, mag string) bool{
	magMap, noteMap := wordMap(mag), wordMap(note)
	for letter, count := range noteMap{
		if value, ok := magMap[letter]; !ok || value <count{
			return false
	}
}
	return true

}

func uniqueOccurences(arr []int)bool{
	counts := map[int]bool{}
	digits := map[int]int{}
	for _, num := range arr{
		digits[num]++
	}

	for _, v := range digits{
		if _, ok:= counts[v]; ok{
			return false
		}else{
			counts[v]= true
		}
	}
	return true
}

func main() {

	mag := "baaab"
	note := "ab"
	// fmt.Println(count(word))
	// fmt.Println(count(word2))

	// fmt.Println(closeStrings(word,word2))
	nums := []int{1,2}
	fmt.Println(wordMap(note))
	fmt.Println(wordMap(mag))
	fmt.Println(RansomeNote(note, mag))
	fmt.Println(uniqueOccurences(nums))

}