func search(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l+r)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid+1
		}
	}
	return -1
}
