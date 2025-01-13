package DynamicProgramming2D

import (
	"math"
)

func MaxProfit(prices []int) int {

	sold, hold, rest := 0, math.MinInt32, 0

	for i := 0; i < len(prices); i++ {

		prevSold := sold
		sold = hold + prices[i]
		hold = max(hold, rest-prices[i])
		rest = max(rest, prevSold)
	}

	return max(sold, rest)
}
