func productExceptSelf(nums []int) []int {
	// products := make([]int, len(nums))
	suffProd := make([]int, len(nums))
	prefProd := make([]int, len(nums))
	prefProd[0] = 1
	suffProd[len(nums)-1] = 1
	for i := 1; i<len(nums);i++ {
		prefProd[i] = nums[i-1] * prefProd[i-1]
	}
	for i := len(nums)-2;i >=0; i-- {
		suffProd[i] = suffProd[i+1] * nums[i+1]
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = prefProd[i] * suffProd[i]
	}
	return res
}
