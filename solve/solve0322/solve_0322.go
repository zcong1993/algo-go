package solve0322

import "math"

/**
 * @index 322
 * @title 零钱兑换
 * @difficulty 中等
 * @tags breadth-first-search,array,dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/coin-change/
 * @frontendId 322
 */

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func CoinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount < 0 {
		return -1
	}

	var coinChange2 func(am int) int
	coinChange2 = func(am int) int {
		if am == 0 {
			return 0
		}

		if am < 0 {
			return -1
		}

		res := math.MaxInt32

		for _, c := range coins {
			sub := coinChange2(am - c)
			if sub == -1 {
				continue
			}
			res = min(res, sub+1)
		}
		if res == math.MaxInt32 {
			return -1
		}
		return res
	}

	return coinChange2(amount)
}

func CoinChange2(coins []int, amount int) int {
	if len(coins) == 0 || amount < 0 {
		return -1
	}

	cache := make(map[int]int)

	var coinChange2 func(am int) int
	coinChange2 = func(am int) int {
		if c, ok := cache[am]; ok {
			return c
		}

		if am == 0 {
			cache[am] = 0
			return 0
		}

		if am < 0 {
			cache[am] = -1
			return -1
		}

		res := math.MaxInt32

		for _, c := range coins {
			sub := coinChange2(am - c)
			if sub == -1 {
				continue
			}
			res = min(res, sub+1)
		}

		if res == math.MaxInt32 {
			cache[am] = -1
			return -1
		}

		cache[am] = res
		return res
	}

	return coinChange2(amount)
}

func CoinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, c := range coins {
			if i-c < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-c])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
