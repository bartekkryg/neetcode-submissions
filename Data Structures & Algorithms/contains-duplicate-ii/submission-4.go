func containsNearbyDuplicate(nums []int, k int) bool {
	exists := make(map[int]bool)
	L := 0
	for R := range nums {
		if R-L > k {
			delete(exists, nums[L])
			L++
		}
		ok := exists[nums[R]]
		if ok {
			return true
		}
		exists[nums[R]] = true
	}
	return false
}
