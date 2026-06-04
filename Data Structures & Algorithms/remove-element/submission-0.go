func removeElement(nums []int, val int) int {
	var toRemove []int 
	var toLeave []int
	for i := range nums {
		if nums[i] == val {
		toRemove = append(toRemove, i)
		} else {
			toLeave = append(toLeave, i)
		}
	}
	var conc []int
	conc = append(toLeave, toRemove...)
	for i := 0; i < len(conc); i++ {
		nums[i] = nums[conc[i]]
	}
	return len(toLeave)
}
