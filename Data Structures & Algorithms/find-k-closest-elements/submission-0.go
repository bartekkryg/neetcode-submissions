func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l+r) / 2
		if x < arr[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	l = l - 1
	r = l + 1
	for r-l-1 < k {
		if l < 0 {
			r++
		} else if r >= len(arr) {
			l--
		} else if abs(arr[l]-x) <= abs(arr[r]-x) {
			l--
		} else {
			r++
		}
	}

	return arr[l+1:r]

}
func abs(a int) int{
	if a < 0 {
		return -a
	}
	return a
}