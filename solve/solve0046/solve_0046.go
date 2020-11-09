package solve0046

/**
@index 46
@title 全排列
@difficulty 中等
@tags backtracking
@draft false
@link https://leetcode-cn.com/problems/permutations/
@frontendId 46
*/

// 回溯
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	// 标记元素是否用过
	usedMap := make(map[int]bool, 0)

	var backtrack func(ans []int)
	backtrack = func(ans []int) {
		// 长度等于数组长度时, 保存结果
		// deep copy
		if len(ans) == len(nums) {
			tmp := make([]int, len(ans))
			copy(tmp, ans)
			res = append(res, tmp)
			return
		}
		// 遍历可选值
		for _, val := range nums {
			// 如果元素已使用, 跳过
			if usedMap[val] {
				continue
			}
			// 选择
			ans = append(ans, val)
			usedMap[val] = true
			// 回溯
			backtrack(ans)
			// 撤销
			usedMap[val] = false
			ans = ans[:len(ans)-1]
		}
	}
	backtrack(make([]int, 0))
	return res
}

func Permute(nums []int) [][]int {
	return permute(nums)
}
