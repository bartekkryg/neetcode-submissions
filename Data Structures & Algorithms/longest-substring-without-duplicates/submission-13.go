func lengthOfLongestSubstring(s string) int {
	uniqueChars := make(map[byte]int)
	L, R := 0, 0
	max := 1
	if len(s) == 0 {
		return 0
	}
	for R < len(s) {
		v, ok := uniqueChars[s[R]]
		if ok {
			if v >= L {
				L=v+1
			}
		} 
		if R - L + 1 > max {
			max = R-L + 1
		}
		uniqueChars[s[R]] = R
		R++
		
	}
	return max
}
