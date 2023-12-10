package top_interview_150

// You can access to this problem via this link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	min := prices[0]
	newMin := prices[0]
	max := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < min && prices[i] < newMin {
			newMin = prices[i]
		}

		if newMin < min && prices[i]-newMin >= max-min {
			max = prices[i]
			min = newMin
		} else if prices[i] > max {
			max = prices[i]
		}
	}

	return max - min
}
