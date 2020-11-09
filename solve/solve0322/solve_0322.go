package solve0322

import "github.com/zcong1993/algo-go/pkg/helper"

/**
@index 322
@title 零钱兑换
@difficulty 中等
@tags dynamic-programming
@draft false
@link https://leetcode-cn.com/problems/coin-change/
@frontendId 322
*/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// dp[i] 表示凑出 i 面值需要最少硬币数
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		// 初始值, amount +1 因为凑够 amount 最多 amount 个
		// 所以 amount + = +inf
		dp[i] = amount + 1
	}
	// base case 0 面值需要 0
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		// dp[i] = min(选择每个面值 + dp[剩下的面值])
		for _, coin := range coins {
			// 硬币面值超出当前 amount
			if i-coin < 0 {
				continue
			}
			dp[i] = helper.Min2(dp[i], 1+dp[i-coin])
		}
	}
	// 如果等于初始值, 表示凑不出来
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func CoinChange(coins []int, amount int) int {
	return coinChange(coins, amount)
}
