package solve0001

/**
 * @index 1
 * @title 两数之和
 * @difficulty 简单
 * @tags array,hash-table
 * @draft false
 * @link https://leetcode-cn.com/problems/two-sum/
 * @frontendId 1
 */

func TwoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if idx, ok := mp[target-cur]; ok {
			return []int{idx, i}
		}
		mp[cur] = i
	}

	return []int{}
}
