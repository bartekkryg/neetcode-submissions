func lengthOfLongestSubstring(s string) int {
	uniqueChars := make(map[byte]int)
	L, R := 0, 1
	max := 1
	currentLength := 1
	if len(s) == 0 {
		return 0
	}
	uniqueChars[s[L]] = L
	for R < len(s) {
		v, ok := uniqueChars[s[R]]
		if ok {
			if currentLength> max {
				max = currentLength
			}
			// delete(uniqueChars, s[L])
			L=v+1
			R=L+1
			uniqueChars = make(map[byte]int)
			uniqueChars[s[L]] = L
			currentLength = 1
		} else {
			uniqueChars[s[R]] = R
			R++
			currentLength++
		}
	}
	if currentLength > max {
		return currentLength
	} else {
		return max
	}
}
