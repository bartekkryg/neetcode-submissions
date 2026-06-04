func longestCommonPrefix(strs []string) string {
    var pref = make([]rune,0)
    main:
    for i := 0; i < 200; i++ {
        if i == len(strs[0]) {
            break main
        }
        ch := []rune(strs[0])[i] 
        for _, str := range strs[1:] {
            if str == "" { 
                return ""
            }
            if i == len(str)||[]rune(str)[i] != ch { 
                break main
            }
        }
        pref = append(pref, ch)
    }
    return string(pref)
}
