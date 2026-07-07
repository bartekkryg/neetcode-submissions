func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	min := 1001
	for l <= r {
		mid := l + (r - l)/2
		if nums[l] <= nums[mid] && nums[mid] <= nums[r] {
			return nums[l]
		} else if nums[l] >= nums[mid] && nums[mid] >= nums[r] {
			return nums[r]
		}
		if nums[l] <= nums[mid] && nums[mid] >= nums[r] {
			l = mid+1
		} else if nums[l] >= nums[mid] && nums[mid] <= nums[r] {
			r = mid
		}
		if nums[mid] <= min {
			min = nums[mid]
		}
	}
	return min
}
