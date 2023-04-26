package blackjack

import (
	"testing"
)

func TestParseCardNotFound(t *testing.T) {

	got := ParseCard("twenty")
	expected := 0

	if got != expected {
		t.Errorf("got %q, wanted %q", got, expected)
	}
}
func TestParseCard(t *testing.T) {
	cases := []struct {
		card     string
		expected int
	}{
		{
			card:     "nine",
			expected: 9,
		},
		{
			card:     "ace",
			expected: 11,
		},
		{
			card:     "king",
			expected: 10,
		},
		{
			card:     "two",
			expected: 2,
		},
	}
	for _, tc := range cases {
		t.Run(tc.card, func(t *testing.T) {
			result := ParseCard(tc.card)
			if result != tc.expected {
				t.Errorf(" expected %q, but got %q", tc.expected, result)
			}
		})
	}
}

func TestFirstTurnSplit(t *testing.T) {
	got := FirstTurn("ace", "ace", "one")
	expected := "P"
	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}

}

func TestFirstTurnStand(t *testing.T) {
	cases := []struct {
		card1    string
		card2    string
		dealer   string
		expected string
	}{
		{
			card1:    "ace",
			card2:    "king",
			dealer:   "ace",
			expected: "S",
		},
		{
			card1:    "ace",
			card2:    "ten",
			dealer:   "ten",
			expected: "S",
		},
		{
			card1:    "ace",
			card2:    "queen",
			dealer:   "ace",
			expected: "S",
		},
		{
			card1:    "ten",
			card2:    "seven",
			dealer:   "eight",
			expected: "S",
		},
		{
			card1:    "four",
			card2:    "eight",
			dealer:   "six",
			expected: "S",
		},
	}
	for _, tc := range cases {
		t.Run(tc.card1, func(t *testing.T) {
			result := FirstTurn(tc.card1, tc.card2, tc.dealer)
			if result != tc.expected {
				t.Errorf("expected %q =, but got %q", tc.expected, result)
			}
		})
	}
}

func TestFirstTurnWin(t *testing.T) {
	cases := []struct {
		card1    string
		card2    string
		dealer   string
		expected string
	}{
		{
			card1:    "ace",
			card2:    "ten",
			dealer:   "two",
			expected: "W",
		},
		{
			card1:    "ace",
			card2:    "jack",
			dealer:   "six",
			expected: "W",
		},
	}
	for _, tc := range cases {
		t.Run(tc.card1, func(t *testing.T) {
			result := FirstTurn(tc.card1, tc.card2, tc.dealer)
			if result != tc.expected {
				t.Errorf("expected %q =, but got %q", tc.expected, result)
			}

		})
	}

}

func TestFirstTurnHit(t *testing.T) {
	cases := []struct {
		card1    string
		card2    string
		dealer   string
		expected string
	}{
		{
			card1:    "seven",
			card2:    "five",
			dealer:   "ten",
			expected: "H",
		},
		{
			card1:    "eight",
			card2:    "eight",
			dealer:   "eight",
			expected: "H",
		},
	}
	for _, tc := range cases {
		t.Run(tc.card1, func(t *testing.T) {
			result := FirstTurn(tc.card1, tc.card2, tc.dealer)
			if result != tc.expected {
				t.Errorf("expected %q =, but got %q", tc.expected, result)
			}
		})

	}

}


