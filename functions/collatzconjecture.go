package main

import (
    "errors"
    "fmt"
)

func CollatzConjecture(n int) (int, error) {
    steps := 0
    if n < 1{
        return steps, errors.New("not possible")
    }
    
    for ; n != 1 ; steps++{
        if n %2 == 0{
            n = n /2
        }else{
        	n = (3*n) + 1
        }
    }
	return steps, nil
}

func main() {
	results, err := CollatzConjecture(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Steps:", results)
	}
}