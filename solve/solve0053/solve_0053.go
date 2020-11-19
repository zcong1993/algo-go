package solve0053

import . "github.com/zcong1993/algo-go/pkg/helper"

/**
 * @index 53
 * @title 最大子序和
 * @difficulty 简单
 * @tags array,divide-and-conquer,dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/maximum-subarray/
 * @frontendId 53
 */

func maxSubArray(nums []int) int {
	sum := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev < 0 {
			prev = nums[i]
		} else {
			prev += nums[i]
		}
		sum = Max2(sum, prev)
	}
	return sum
}

func MaxSubArray(nums []int) int {
	return maxSubArray(nums)
}
