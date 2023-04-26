package blackjack

var cards = map[string]int{
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"ace":   11,
}

func ParseCard(card string) int {
	for key, value := range cards {
		if card == key {
			return value
		}
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardTotal := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case cardTotal == 22:
		return "P"
	case cardTotal == 21:
		if dealer == 10 || dealer == 11 {
			return "S"
		} else {
			return "W"
		}
	case cardTotal >= 17 && cardTotal <= 20:
		return "S"
	case cardTotal >= 12 && cardTotal <= 16:
		if dealer >= 7 {
			return "H"
		} else {
			return "S"
		}
	default:
		return "H"
	}
}