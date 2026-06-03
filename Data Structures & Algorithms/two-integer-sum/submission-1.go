func twoSum(nums []int, target int) []int {
    var hash = make(map[int]int)
    for i := range nums {
        diff := target - nums[i]
        j, ok := hash[diff]
        if !ok {
            hash[nums[i]] = i
            continue
        }
        return []int{j, i}
    }
    return nil

}
