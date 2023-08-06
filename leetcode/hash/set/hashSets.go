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

func main() {

	word := "ab"
	word2 := "aa"
	fmt.Println(count(word))
	fmt.Println(count(word2))

	fmt.Println(closeStrings(word,word2))

}