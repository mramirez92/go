package main

import (
	"fmt"
)


func stringCompression(chars []byte)(int, []byte){
	group := []byte{}
	for i:= 0; i <len(chars);{
		current, counts := chars[i], 1

		for j := i+1; j <len(chars) && chars[j] == current; j++{
			counts ++
		}
		group = append(group, current)
		if counts >1{
				group = append(group, '0'+byte(counts))
			}
		chars = append(chars[:i], append(group, chars[i+counts:]...)...)
		counts, group = 1, []byte{}
		i += counts
	}
	return len(chars), chars
}

func main() {
	b := []byte{'b','b','b','b','b','b','b','b','b','a','a'}

	// ba := []byte{'b','b','b','a','a','a','a','c','c'}


	_ , arr := stringCompression(b)
	fmt.Println(arr)
	// fmt.Println(string(arr[len(arr)-1]))

	for _, char := range arr{
		fmt.Println(string(char))
	}
}

