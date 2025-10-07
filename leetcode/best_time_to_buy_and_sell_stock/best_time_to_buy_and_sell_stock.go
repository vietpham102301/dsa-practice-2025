package main

import "fmt"

func calProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	smlIndex := 0
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[smlIndex] {
			smlIndex = i
		}

		if i != smlIndex && prices[i]-prices[smlIndex] > maxProfit {
			maxProfit = prices[i] - prices[smlIndex]
		}

	}

	return maxProfit
}

func main() {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{2, 4, 1}, 2},
		{[]int{5}, 0},
		{[]int{}, 0},
		{[]int{3, 3, 3, 3}, 0},
		{[]int{1, 2, 100}, 99},
	}

	for i, tc := range tests {
		res := calProfit(tc.prices)
		status := "PASS"
		if res != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("Test %02d: %s\n", i+1, status)
		fmt.Printf("  Prices   : %v\n", tc.prices)
		fmt.Printf("  Expected : %d\n", tc.expected)
		fmt.Printf("  Result   : %d\n\n", res)
	}
}
