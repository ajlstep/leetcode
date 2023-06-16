package besttimetobuy

import "math"

func maxProfit(prices []int) int {
	minimal := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > minimal {
				minimal = prices[j] - prices[i]
			}
		}
	}
	return minimal
}

func maxProfitv2(prices []int) int {
	// It is impossible to have stock to sell on first day, so -infinity is set as initial value
	cur_hold, cur_not_hold := math.MinInt64, 0

	for _, stock_price := range prices {

		prev_hold, prev_not_hold := cur_hold, cur_not_hold

		// Either keep in hold, or just buy today with stock price
		cur_hold = Max(prev_hold, 0-stock_price)

		// Either keep in not hold, or just sell stock with stock price
		cur_not_hold = Max(prev_not_hold, prev_hold+stock_price)
	}

	// Max profit must be in not_hold state finally
	return cur_not_hold

}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
