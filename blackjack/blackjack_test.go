package blackjack

import "testing" 

func TestParseCardNotFound(t *testing.T){

	got := ParseCard("twenty")
	want := -1

	if got != want{
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestParseCard(t *testing.T){
	cases := []struct{
		card string
		expected int
	}{
		{
			card: "nine",
			expected: 9 ,
		},
		{
			card: "ace", 
			expected: 11,

		},
		{
			card: "king",
			expected: 10,

		},
		{
			card: "two", 
			expected: 2,
		},
	}
	for _, tc := range cases {
		t.Run(tc.card, func(t *testing.T){
			result := ParseCard(tc.card)
			if result != tc.expected{
				t.Errorf(" expected %q, but got %q", tc.expected, result)
			}
		})
	}
}

func TestFirstTurnSplit(t *testing.T) {
	cases := []struct{
		card1 string
		card2 string
		dealer string 
		expected string 
	}{
		{
			card1: "ace",
			card2: "king",
			dealer: "one",
			expected: "W",
		},
		{
			card1: "ace",
			card2: "ten",
			dealer: "ten",
			expected: "S",
		},
		{
			card1: "ace",
			card2: "queen",
			dealer: "ace",
			expected: "S",
		},
	}
	for _, tc := range cases{
		t.Run(tc.card1, func(t *testing.T) {
			result := FirstTurn(tc.card1, tc.card2, tc.dealer)
			if result != tc.expected{
				t.Errorf("expected %q =, but got %q", tc.expected, result)
			}
		})
	}
}



