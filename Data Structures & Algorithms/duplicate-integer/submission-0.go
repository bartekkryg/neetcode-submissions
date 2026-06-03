func hasDuplicate(nums []int) bool {
    for i := range nums {
        for _, k := range nums[i+1:] {
            if nums[i] == k {
                return true
            }
        }
    }
    return false
}
