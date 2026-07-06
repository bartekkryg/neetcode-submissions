func rotate(nums []int, k int) {
	var nums2 = make([]int, len(nums))
	copy(nums2, nums)
	r := len(nums)-1
	k %= len(nums)
	for i := 0; i < k; i++ {
		nums2 = append([]int{nums2[r]}, nums2[0:r]...)
	}
	copy(nums, nums2)
}
