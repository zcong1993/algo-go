package solve0300

import . "github.com/zcong1993/algo-go/pkg/helper"

/**
 * @index 300
 * @title 最长递增子序列
 * @difficulty 中等
 * @tags array,binary-search,dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/longest-increasing-subsequence/
 * @frontendId 300
 */

func LengthOfLIS(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return res
	}

	// dp[i] 指的是 [0..i] 段的最长递增子序列结果
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		// 搜索 [0, i] 段
		for j := 0; j < i; j++ {
			// i > j 才会和 i 构成递增序列
			if nums[i] > nums[j] {
				// dp[j] 加上 i 自己
				dp[i] = Max2(dp[i], dp[j]+1)
			}
		}
		res = Max2(res, dp[i])
	}

	return res
}
