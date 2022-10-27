package solve0198

import . "github.com/zcong1993/algo-go/pkg/helper"

/*
*
@index 198
@title 打家劫舍
@difficulty 简单
@tags dynamic-programming
@draft false
@link https://leetcode-cn.com/problems/house-robber/
@frontendId 198.
*/
func rob(nums []int) int {
	// 从 start 开始打劫最大利润
	var dp func(start int) int
	dp = func(start int) int {
		// 越界收益为 0
		if start >= len(nums) {
			return 0
		}
		// 1. 第 i 家不打劫, 收益为 dp(start+1) 从下一家开始
		// 2. 第 i 家打劫, 收益为 nums[start] + dp(start+2), 从第二家开始
		return Max2(dp(start+1), nums[start]+dp(start+2))
	}
	return dp(0)
}

func robDp(nums []int) int {
	dp := make([]int, len(nums)+2)
	dp[len(nums)] = 0
	dp[len(nums)+1] = 0
	// 可以看出 i 状态只和 i+1 和 i+2 有关
	// 可以优化成常数空间
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = Max2(dp[i+1], dp[i+2]+nums[i])
	}
	return dp[0]
}

func robDpOptim(nums []int) int {
	var dp int
	dpI1, dpI2 := 0, 0
	// 可以看出 i 状态只和 i+1 和 i+2 有关
	// 可以优化成常数空间
	for i := len(nums) - 1; i >= 0; i-- {
		dp = Max2(dpI1, dpI2+nums[i])
		dpI1, dpI2 = dp, dpI1
	}
	return dp
}

func Rob(nums []int) int {
	return rob(nums)
}

func RobDp(nums []int) int {
	return robDp(nums)
}

func RobDpOptim(nums []int) int {
	return robDpOptim(nums)
}
