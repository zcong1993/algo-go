package solve0001

/**
@index 1
@title 两数之和
@difficulty 简单
@tags array,hash-table
@draft false
@link https://leetcode-cn.com/problems/two-sum
*/
func TwoSum(nums []int, target int) []int {
	mm := make(map[int]int, len(nums))
	for idx, val := range nums {
		if i, ok := mm[target-val]; ok {
			return []int{i, idx}
		}
		mm[val] = idx
	}
	return []int{}
}
