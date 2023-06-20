package space

import (
	"math"
	"testing"
)

func TestSpace(t *testing.T) {

	cases := []struct {
		planet   string
		seconds  float64
		expected float64
	}{{
		planet:   "Mercury",
		seconds:  2134835688.0,
		expected: 280.88,
	},
		{
			planet:   "Sun",
			seconds:  68080999.0,
			expected: -1.00,
		},
		{
			planet:   "Earth",
			seconds:  220903200,
			expected: 7.00,
		},
	}
	for _, tc := range cases {
		result := Age(tc.seconds, Planet(tc.planet))
		roundedResult := math.Round(result*100) / 100
		roundedExpected := math.Round(tc.expected*100) / 100

		if roundedResult != roundedExpected {
			t.Errorf("For planet %s, got %.2f, expected %.2f", tc.planet, roundedResult, roundedExpected)
		}
	}
}
