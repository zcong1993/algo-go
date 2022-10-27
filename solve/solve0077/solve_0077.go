package solve0077

/*
*
@index 77
@title 组合
@difficulty 中等
@tags backtracking
@draft false
@link https://leetcode-cn.com/problems/combinations/
@frontendId 77.
*/
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0)
	var backtrack func(s int)
	backtrack = func(s int) {
		// 剪枝优化
		// 当前结果长度 + 剩下的元素个数 < k 时, 是无解的, 提前退出
		if len(ans)+(n-s+1) < k {
			return
		}
		// 如果答案长度 == k, 保存结果
		// 注意 deep copy
		if len(ans) == k {
			tmp := make([]int, len(ans))
			copy(tmp, ans)
			res = append(res, tmp)
			return
		}
		// 从当前开始, 避免重复
		for i := s; i <= n; i++ {
			ans = append(ans, i)
			backtrack(i + 1)
			ans = ans[:len(ans)-1]
		}
	}
	backtrack(1)
	return res
}

func Combine(n int, k int) [][]int {
	return combine(n, k)
}
