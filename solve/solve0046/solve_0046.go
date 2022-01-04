package solve0046

/**
 * @index 46
 * @title 全排列
 * @difficulty 中等
 * @tags array,backtracking
 * @draft false
 * @link https://leetcode-cn.com/problems/permutations/
 * @frontendId 46
 */

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func Permute(nums []int) [][]int {
	l := len(nums)
	var res [][]int
	var backtrack func(track []int)
	backtrack = func(track []int) {
		if len(track) == l {
			// 注意 copy slice
			tmp := make([]int, l)
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for _, n := range nums {
			if contains(track, n) {
				continue
			}
			// 选择
			track = append(track, n)
			backtrack(track)
			// 撤销
			track = track[:len(track)-1]
		}
	}
	backtrack([]int{})
	return res
}
