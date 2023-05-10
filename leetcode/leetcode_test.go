package leetcode

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