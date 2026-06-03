func hasDuplicate(nums []int) bool {
    var values = make(map[int]bool)
    for _, v := range nums {
        exists := values[v]
        if !exists {
            values[v] = true
            continue
        }
        return true
    }
    return false
}
