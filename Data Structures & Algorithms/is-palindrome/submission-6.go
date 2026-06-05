func isPalindrome(s string) bool {
	t := strings.ToLower(s)

	var result strings.Builder
    for i := 0; i < len(t); i++ {
        b := t[i]
        if ('a' <= b && b <= 'z') || ('0' <= b && b <= '9')  {
            result.WriteByte(b)
        }
    }
    t = result.String()
	
	for i := 0; i < len(t)/2; i++ {
		end := len(t)-1-i
		if t[i] != t[end] {
			return false
		}
	}
	return true
}
