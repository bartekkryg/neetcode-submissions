func maxProfit(prices []int) int {
	L := 0
	R := 1
	maxProfit := 0
	for R < len(prices) {
		profit := prices[R] - prices[L]
		if profit > maxProfit {
			maxProfit = profit
		} 
		if prices[L] >= prices[R] {
			L = R
		}
		R++
		
	}
	if maxProfit <= 0 {
		return 0
	}
	return maxProfit
}
