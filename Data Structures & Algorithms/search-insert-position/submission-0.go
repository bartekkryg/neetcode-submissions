func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if nums[l] > target {
		return 0
	} else if nums[r] < target {
		return len(nums)
	}
	for l < r {
		mid := (l+r)/2
		if target > nums[mid] {
			l = mid+1
		} else {
			r = mid
		}
	}
	return l
}