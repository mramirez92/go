package main

import (
	"strings"
	"fmt"
)


func ScrabbleScore(word string)(total int){
	for _, char := range strings.ToUpper(word){
		switch char{
		case 'Q' , 'Z':
			total += 10
		case 'J', 'X':
			total += 8
		case 'K':
			total += 5
		case 'F', 'H', 'V', 'W', 'Y':
			total += 4
		case 'B', 'C', 'M', 'P':
			total+=3
		case 'D', 'G':
			total +=2
		default:
			total ++
		}	
	}
	return total
}

func main(){
	str := "aei"
	fmt.Println(ScrabbleScore(str))

}