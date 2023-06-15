package slices

import "testing"

func TestGetItem(t *testing.T) {
	cases := []struct {
		slice    []int
		index    int
		expected int
	}{
		{
			slice:    []int{0, 1, 2, 3},
			index:    3,
			expected: 3,
		},
		{
			slice:    []int{0, 1, 2, 3, 4},
			index:    -2,
			expected: -1,
		},
		{
			slice:    []int{0, 1, 2, 3, 4},
			index:    5,
			expected: -1,
		},
	}
	for _, tc := range cases {
		results := GetItem(tc.slice, tc.index)
		if results != tc.expected {
			t.Errorf("got %+v, expected %+v", results, tc.expected)
		}
	}
}
