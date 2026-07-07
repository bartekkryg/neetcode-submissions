func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	min := nums[l]
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			r=mid
		} else  {
			l=mid+1
		}
		if nums[l] <= min {
			min = nums[l]
		}
	}
	return min
}
