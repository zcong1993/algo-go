package solve0031

/**
 * @index 31
 * @title 下一个排列
 * @difficulty 中等
 * @tags array,two-pointers
 * @draft false
 * @link https://leetcode-cn.com/problems/next-permutation/
 * @frontendId 31
 */

func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	l := len(nums)
	i, j, k := l-2, l-1, l-1
	// 从右到左找到第一个升序 i,j, 保证 [j, end) 降序
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		// 从右到左找到第一个大于 i 的数
		for nums[i] >= nums[k] {
			k--
		}
		// 交换 i k
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 反转数组 [j, end], 变成升序会更小
	for i, j := j, l-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
