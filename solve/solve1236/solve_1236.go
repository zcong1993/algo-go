package solve1236

/*
*
@index 1236
@title 第 N 个泰波那契数
@difficulty 简单
@tags recursion
@draft false
@link https://leetcode-cn.com/problems/n-th-tribonacci-number/
@frontendId 1137.
*/
func Tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	prev3, prev2, prev1 := 0, 1, 1
	for i := 3; i <= n; i++ {
		prev3, prev2, prev1 = prev2, prev1, prev1+prev2+prev3
	}
	return prev1
}
