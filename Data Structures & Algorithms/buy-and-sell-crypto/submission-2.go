func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for right := 1; right < len(prices); right++ {
		profit := prices[right] - minPrice

		if profit > maxProfit {
			maxProfit = profit
		}

		if prices[right] < minPrice {
			minPrice = prices[right]
		}
	}

	return maxProfit
}
