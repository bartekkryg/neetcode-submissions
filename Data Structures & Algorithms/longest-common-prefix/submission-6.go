func longestCommonPrefix(strs []string) string {
    pref := []rune(strs[0])
    for _, str := range strs[1:] {
        if str == "" {
            return ""
        }
        for i, ch := range str {
            if len(str) < len(pref) {
                pref = pref[:len(str)]
            }
            if i < len(pref) && pref[i] != ch {
                pref = pref[:i]
            }

        }
    }
    return string(pref)
}
