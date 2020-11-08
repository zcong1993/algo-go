package solve0122

import "github.com/zcong1993/algo-go/pkg/helper"

/**
@index 122
@title 买卖股票的最佳时机 II
@difficulty 简单
@tags greedy,array
@draft false
@link https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
@frontendId 122
*/

// 贪心算法
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

// 动态规划
// dp[i][1 or 0] i 表示第几天的最大收益, 1 代表手里有股票, 0 代表没有
// dp[i][1] i 天手里有股票: 1. 上一天手里有股票, 今天啥都不干 2. 昨天没有, 今天买了
// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
// dp[i][0] i 天手里没股票: 1. 上一天没股票, 今天啥都不干 2. 昨天买了, 今天卖了
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
// base case dp[0][0] = 0 没操作 dp[0][1] = -prices[0] 第一天买了
func maxProfitDp(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [2]int{0, 0}
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	// 可以看到, dp[i] 只依赖于 dp[i-1]
	// 所以可以优化成常数空间复杂度, 和斐波那契一样
	for i := 1; i < n; i++ {
		dp[i][0] = helper.Max2(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = helper.Max2(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func maxProfitDpOptim(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dpPrev0, dpPrev1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dpPrev0 = helper.Max2(dpPrev0, dpPrev1+prices[i])
		dpPrev1 = helper.Max2(dpPrev1, dpPrev0-prices[i])
	}
	return dpPrev0
}

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}

func MaxProfitDp(prices []int) int {
	return maxProfitDp(prices)
}

func MaxProfitDpOptim(prices []int) int {
	return maxProfitDpOptim(prices)
}
