package blackjack
import (
	"fmt"
)

 var cards = map[string]int{
	"two": 2,
	"three":3,
	"four": 4, 
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8, 
	"nine": 9,
	"ten" :10,
	"jack": 10,
	"queen": 10,
	"king": 10,
	"ace": 11,
}

func ParseCard(card string) int {
    for key, value := range cards {
        if card == key {
            return value
        }
    }
    fmt.Println("not found")
    return -1
}



// func main(){
// 	fmt.Println("Enter card:")
// 	var userInput string
// 	fmt.Scanln(&userInput)
// 	fmt.Println(ParseCard(strings.ToLower(userInput)))

	
// 	}
