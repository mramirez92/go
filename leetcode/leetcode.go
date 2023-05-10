package leetcode

func IsSubsequence(s, t string) bool{
	i := 0
	for _, char  := range t{
		if i < len(s) &&  byte(char) == s[i]{
			i++	
		}
	}
	return i == len(s)
}