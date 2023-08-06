package main

import "testing"

func TestIsSub(t *testing.T){
	cases := [] struct {
		subsequence string
		testString string
		expected bool
	}{
		{
			subsequence: "abc",
			testString: "anbnc",
			expected: true,
		},
		{
			subsequence: "aa",
			testString: "bb",
			expected: false,
		},
		{
			subsequence: "monica",
			testString: "zmyxonwica",
			expected: true,
		},
	}
	for _, tc := range cases{
		results := IsSubsequence(tc.subsequence,tc.testString)
		if results != tc.expected{
			t.Errorf("got%+v, expected %+v", results, tc.expected)
		}
	}
}

func TestFindPivot(t *testing.T) {
	cases := [] struct {
		nums []int
		expected int
	}{
		{
			nums: []int {5, 2, 3, 4, 3}, 
			expected: 2,
		},
		{
			nums: []int {2, 1, -1}, 
			expected: 0,
		},
		{
			nums: []int {2,7},
			expected: -1,
		},
	}
	for _, tc := range cases{
		results := FindPivot(tc.nums)
		if results != tc.expected { 
			t.Errorf("got %+v, expected %+v", results, tc.expected)
		}
	}

}