package solve0041

/**
 * @index 41
 * @title 缺失的第一个正数
 * @difficulty 困难
 * @tags array,hash-table
 * @draft false
 * @link https://leetcode-cn.com/problems/first-missing-positive/
 * @frontendId 41
 */

func FirstMissingPositive(nums []int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for nums[i] > 0 && nums[i] <= l && nums[nums[i]-1] != nums[i] {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	for i := 0; i < l; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return l + 1
}
