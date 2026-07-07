func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if target > nums[mid] {
			l = mid+1
		} else {
			r = mid
		}
	}
	return l
}