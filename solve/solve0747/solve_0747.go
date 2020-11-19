package solve0747

import (
	. "github.com/zcong1993/algo-go/pkg/helper"
)

/**
 * @index 747
 * @title 使用最小花费爬楼梯
 * @difficulty 简单
 * @tags array,dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/min-cost-climbing-stairs/
 * @frontendId 746
 */

// https://leetcode-cn.com/problems/min-cost-climbing-stairs/solution/yi-bu-yi-bu-tui-dao-dong-tai-gui-hua-de-duo-chong-/
func minCostClimbingStairs(cost []int) int {
	// dp[i] 表示踏上 i 阶需要的最小体力
	dp := make([]int, len(cost))
	dp[0] = cost[0] // 只能踩第一阶
	dp[1] = cost[1] // 越过第一阶
	for i := 2; i < len(cost); i++ {
		// 到达第 i 层, 要么是 i-1, 要么是 i-2 + cost[i]
		dp[i] = Min2(dp[i-1], dp[i-2]) + cost[i]
	}
	// 结果为, min(i-2, i-1) 可以跨过去, 倒数第二阶(i-2)和倒数第一阶(i-1)
	return Min2(dp[len(cost)-2], dp[len(cost)-1])
}

// 压缩空间
func MinCostClimbingStairs2(cost []int) int {
	prev2, prev1 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		// 到达第 i 层, 要么是 i-1, 要么是 i-2 + cost[i]
		prev2, prev1 = prev1, Min2(prev1, prev2)+cost[i]
	}
	return Min2(prev2, prev1)
}

func MinCostClimbingStairs(cost []int) int {
	return minCostClimbingStairs(cost)
}
