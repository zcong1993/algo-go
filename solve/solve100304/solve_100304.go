package solve100304

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

func MaxSubArray(nums []int) int {
	return maxSubArray(nums)
}
