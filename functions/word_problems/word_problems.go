package main 

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
)


func Filter(problem string) (result []string){
	nReg := regexp.MustCompile(`^-?\d+$`)
	opReg :=  regexp.MustCompile(`^(plus|minus|multiplied|divided)$`)

	for _, ele := range strings.Split(problem, " "){
		if nReg.MatchString(ele) || opReg.MatchString(ele) {
			result = append(result, ele)
		}
	}
	return 
}


func Answer(problem string) (int, bool) {
	elements := Filter(problem)

	if len(elements) == 0 {
		return 0, true
	}

	if len(elements) == 1 {
		result, err := strconv.Atoi(elements[0])
		if err != nil {
			return 0, true
		}
		return result, false
	}

	operators := map[string]func(int, int) int{
		"plus":       func(a, b int) int { return a + b },
		"minus":      func(a, b int) int { return a - b },
		"multiplied": func(a, b int) int { return a * b },
		"divided":    func(a, b int) int { return a / b },
	}

	result, err := strconv.Atoi(elements[0])
	if err != nil {
		return 0, true
	}

	operatorIndex := 1
	for operatorIndex < len(elements) {
		operator := elements[operatorIndex]
		nextNum, err := strconv.Atoi(elements[operatorIndex+1])
		if err != nil {
			return 0, true
		}

		if operation, found := operators[operator]; found {
			result = operation(result, nextNum)
		} else {
			return 0, true
		}

		operatorIndex += 2
	}

	return result, false
}




func main() {
	s := "what is 5 plus 10"
	fmt.Println(Filter(s))
	fmt.Println(Answer(s))

}
