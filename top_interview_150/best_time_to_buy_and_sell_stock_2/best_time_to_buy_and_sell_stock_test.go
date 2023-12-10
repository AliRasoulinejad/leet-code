package best_time_to_buy_and_sell_stock_2

import (
	"testing"
)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	testCases := []struct {
		inputPrices    []int
		expectedProfit int
	}{
		{
			inputPrices:    []int{7, 1, 5, 3, 6, 4},
			expectedProfit: 7,
		},
		{
			inputPrices:    []int{1, 2, 3, 4, 5},
			expectedProfit: 4,
		},
		{
			inputPrices:    []int{7, 6, 4, 3, 1},
			expectedProfit: 0,
		},
		{
			inputPrices:    []int{1, 2},
			expectedProfit: 1,
		},
		{
			inputPrices:    []int{4, 7, 1, 2, 11},
			expectedProfit: 13,
		},
	}

	for _, testCase := range testCases {
		response := maxProfit(testCase.inputPrices)
		if response != testCase.expectedProfit {
			t.Errorf("expected response is: %d, but got %d", testCase.expectedProfit, response)
		}
	}
}
