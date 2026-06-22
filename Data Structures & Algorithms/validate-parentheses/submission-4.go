func isValid(s string) bool {
    stack := make([]rune, 0)
    closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

    for _, c := range s {
        if open, exists := closeToOpen[c]; exists {
            if len(stack) > 0 {
				top := stack[len(stack)-1]
                if top != open {
                    return false
                }
				stack = stack[:len(stack)-1]
            } else {
                return false
            }
        } else {
            stack = append(stack, c)
        }
    }

    return len(stack) == 0 
}