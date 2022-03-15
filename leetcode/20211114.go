package leetcode

import "math"

func maxProfit(prices []int) int {
	maximalProfit, l, r := 0, 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] >= prices[l] {
			r = i
			if prices[r]-prices[l] > maximalProfit {
				maximalProfit = prices[r] - prices[l]
			}
		} else {
			l = i
			r = i
		}
	}
	return maximalProfit
}

var mapping []int

func maxProfitV2Re(index int, prices []int) int {
	if mapping[index] >= 0 {
		return mapping[index]
	}
	maximalProfit := 0
	for i := index; i < len(prices)-1; i++ {
		maxProfit := math.MinInt32
		for j := i + 1; j < len(prices); j++ {
			profit := (prices[j] - prices[i]) + maxProfitV2Re(j+1, prices)
			if profit > maxProfit {
				maxProfit = profit
			}
		}
		if maxProfit > maximalProfit {
			maximalProfit = maxProfit
		}
	}

	mapping[index] = maximalProfit
	return maximalProfit
}

func maxProfitV2(prices []int) int {
	mapping = make([]int, len(prices))
	for i := 0; i < len(mapping); i++ {
		mapping[i] = -1
	}
	return maxProfitV2Re(0, prices)
}

// 第i天状态为state能够获得的最大利润
func profit(prices []int, i, state int) int {

	if i == len(prices)-1 {
		if state == 1 {
			return prices[i]
		} else {
			return 0
		}
	}

	// 手上没有股票，可以1.继续维持这种状态；2.买入股票
	if state == 0 {
		return int(math.Max(float64(profit(prices, i+1, 0)), float64(profit(prices, i+1, 1)-prices[i])))
	} else if state == 1 { // 手上有股票，可以1.继续维持这种状态；2.卖出股票，进入冰冻
		return int(math.Max(float64(profit(prices, i+1, 1)), float64(profit(prices, i+1, -1)+prices[i])))
	} else { // 处于冰冻期，可以1.继续维持这种状态，到i+1天解冻；这里也可以是i+k天
		return profit(prices, i+1, 0)
	}
}

func maxProfit4(prices []int) int {

	return profit(prices, 0, 0)
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func maxProfit5(prices []int) int {

	// 0:那一天结束“之后”没有股票且不处于冷冻期（前一天结束了本来就没有股票，今天啥也没干/前一天结束处于冷冻期，冷冻了今天）
	// 1:那一天结束“之后”有股票（前一天结束了本来就有股票，今天继续持有/前一天结束了没有股票且不能处于冷冻期，（否则今天就是冷冻期），今天买入的）
	// 2:那一天结束“之后”没有股票且处于冷冻期（前一天结束了股票必须有，今天卖）
	dp := make([][3]int, len(prices))
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}

	return max(dp[len(prices)-1][0], dp[len(prices)-1][2])
}
