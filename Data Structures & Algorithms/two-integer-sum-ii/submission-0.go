func twoSum(numbers []int, target int) []int {
	L := 0
	R := len(numbers)-1
	for L < len(numbers) {
		sum := numbers[L] + numbers[R]
		if sum == target {
			return []int{L+1,R+1}
		} else if sum > target {
			R--
		} else if sum < target {
			L++
		}
	}
	return []int{0,0}
}
