import (
    "maps"
)
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var charsS = make(map[rune]int)
    for _, ch := range s {
        charsS[ch]++
    }
    var charsT = make(map[rune]int)
    for _, ch := range t {
        _, ok := charsS[ch]
        if !ok {
            return false
        }
        charsT[ch]++
    }

    return maps.Equal(charsS, charsT)
    
}
