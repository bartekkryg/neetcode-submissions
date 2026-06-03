func twoSum(nums []int, target int) []int {
    var hash = make(map[int][]int)
    for i := 0; i < len(nums); i++ {
        for k := i+1; k<len(nums); k++ {
            indices := []int{i,k}
            sum := nums[i] + nums[k]
            hash[sum] = indices
        }
    }
    for k, v := range hash {
        if k == target {
            return v
        }
    }
    return nil
}
