package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	min := math.MaxInt
	stonks := 0
	for _, price := range prices {
		current := price - min
		if current > stonks {
			stonks = current
		}
		if price < min {
			min = price
		}
	}
	return stonks
}

func maxProfit2(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func maxProfitFree(prices []int, fee int) int {
	profit := 0
	current := prices[0]
	for i := 1; i < len(prices); i++ {

		if prices[i] < current {
			current = prices[i]
		}
		if prices[i] > current+fee {
			profit += prices[i] - current - fee
			current = prices[i]
		}
	}
	return profit
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func test(prices []int, fee int) int {
	if len(prices) <= 1 {
		return 0
	}
	days := len(prices)
	buy := make([]int, days)
	sell := make([]int, days)
	buy[0] = -prices[0] - fee
	sell[0] = 0
	for i := 1; i < days; i++ {
		buy[i] = max(buy[i-1], sell[i-1]-prices[i]-fee) // keep the same as day i-1, or buy from sell status at day i-1
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])    // keep the same as day i-1, or sell from buy status at day i-1
	}
	return sell[days-1]
}

func main() {
	// fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(test([]int{1, 3, 7, 5, 10, 3}, 0))
}
