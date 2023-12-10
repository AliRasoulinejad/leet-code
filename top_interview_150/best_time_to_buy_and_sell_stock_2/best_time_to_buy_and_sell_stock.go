package best_time_to_buy_and_sell_stock_2

// You can access to this problem via this link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	min := prices[0]
	max := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] >= max {
			max = prices[i]
		} else {
			profit += max - min
			max = prices[i]
			min = prices[i]
		}
	}

	profit += max - min

	return profit
}
