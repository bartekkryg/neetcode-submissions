func removeDuplicates(nums []int) int {
    k := 1
    for R := 1; R < len(nums); R++ {
        if nums[R-1] != nums[R] {
            nums[k] = nums[R]
            k++
        }
    }
    return k
}
