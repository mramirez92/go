package main

import (
	"fmt"
	"strconv"
)

func stringComprehension(chars []byte) (int, []byte) {
	rewriteIdx := 0
	for i := 0; i < len(chars); {
		current, counts := chars[i], 1
		for j := i + 1; j < len(chars) && chars[j] == current; j++ {
			counts++
		}
		chars[rewriteIdx] = current
		rewriteIdx++

		if counts > 1 {
			for _, num := range strconv.Itoa(counts) {
				chars[rewriteIdx] = byte(num)
				rewriteIdx++
			}
		}
		i += counts
	}
	return rewriteIdx, chars
}

func main() {
	b := []byte{'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'a', 'a'}
	// longBoi := []byte{'p','p','p','p','m','m','b','b','b','b','b','u','u','r','r','u','n','n','n','n','n','n','n','n','n','n','n','u','u','u','u','a','a','u','u','r','r','r','s','s','a','a','y','y','y','g','g','g','g','g'}

	// ba := []byte{'b','b','b','a','a','a','a','c','c'}

	num, arr := stringComprehension(b)
	fmt.Println(string(arr[:num]))

}
