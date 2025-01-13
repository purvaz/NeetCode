package SlidingWindow

import "fmt"

func MaxProfit(prices []int) int {

	maximum := 0

	for i, j := 0, 1; j < len(prices) && i < j; {
		fmt.Println(prices[i], prices[j])
		if prices[j]-prices[i] > maximum {
			maximum = prices[j] - prices[i]
		}
		if prices[j] <= prices[i] {
			i = j
			j++
		} else {
			j++
		}
	}
	return maximum
}
