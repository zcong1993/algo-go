package solve100304

import . "github.com/zcong1993/algo-go/pkg/helper"

/**
@index 100304
@title 连续子数组的最大和
@difficulty 简单
@tags divide-and-conquer,dynamic-programming
@draft false
@link https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
@frontendId 剑指 Offer 42
*/

func maxSubArray(nums []int) int {
	maxSum, curSum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		// 当前以 i-1 结尾的和 >= 0 时, 说明对 curSum 贡献为正
		if curSum >= 0 {
			curSum += nums[i]
			// curSum 为负时, curSum + nums[i] < nums[i]
			// 所以 curSum 直接选 nums[i]
		} else {
			curSum = nums[i]
		}

		// 比较是否需要更新结果
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func maxSubArrayDp(nums []int) int {
	maxSum := nums[0]
	// dp[i] 表示以 i 结尾的最大子数组和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// dp[i] = max(加上当前元素, 以当前元素开始忽略之前的元素序列)
		// 可以看到 dp[i] 只和 dp[i-1] 有关, 可以压缩空间
		dp[i] = Max2(dp[i-1]+nums[i], nums[i])
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	return maxSum
}

func maxSubArrayDp2(nums []int) int {
	maxSum := nums[0]
	preSum := nums[0]
	for i := 1; i < len(nums); i++ {
		// dp[i] = max(加上当前元素, 以当前元素开始忽略之前的元素序列)
		// 可以看到 dp[i] 只和 dp[i-1] 有关, 可以压缩空间
		cur := Max2(preSum+nums[i], nums[i])
		if cur > maxSum {
			maxSum = cur
		}
		preSum = cur
	}
	return maxSum
}

func MaxSubArray(nums []int) int {
	return maxSubArray(nums)
}

func MaxSubArrayDp(nums []int) int {
	return maxSubArrayDp(nums)
}

func MaxSubArrayDp2(nums []int) int {
	return maxSubArrayDp2(nums)
}
