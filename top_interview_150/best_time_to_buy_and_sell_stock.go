package top_interview_150

// You can access to this problem via this link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(allPrices []int) int {
	var maxProfit int
	for i := 1; i < len(allPrices); i++ {
		prices := allPrices[i:]

		buy := prices[0]
		sell := prices[0]
		for j := 1; j < len(prices); j++ {
			if prices[j] > sell {
				sell = prices[j]
			}
		}

		profit := sell - buy
		if profit > maxProfit {
			maxProfit = sell - buy
		}
	}

	return maxProfit
}
