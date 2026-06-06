func validPalindrome(s string) bool {
	isPalindrome := func(j, end int) bool {
		for i := j; i < end; i++ {
			if s[i] != s[end] {
				return false
			}
			end--
		}
		return true
	}
	end := len(s)-1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[end] {
			return isPalindrome(i, end-1) || isPalindrome(i+1, end)
		}
		end--
	}
	return true
}
