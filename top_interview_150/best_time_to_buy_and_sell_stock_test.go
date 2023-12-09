package top_interview_150

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
			expectedProfit: 5,
		},
		{
			inputPrices:    []int{7, 6, 4, 3, 1},
			expectedProfit: 0,
		},
	}

	for _, testCase := range testCases {
		response := maxProfit(testCase.inputPrices)
		if response != testCase.expectedProfit {
			t.Errorf("expected response is: %d, but got %d", testCase.expectedProfit, response)
		}
	}
}
