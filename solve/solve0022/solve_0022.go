package solve0022

/*
*
@index 22
@title 括号生成
@difficulty 中等
@tags string,backtracking
@draft false
@link https://leetcode-cn.com/problems/generate-parentheses/
@frontendId 22.
*/
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}
	var backtrack func(ans string, lUsed, rUsed int)
	// 回溯, 附带左右括号已使用数量
	backtrack = func(ans string, lUsed, rUsed int) {
		// 左右已使用数量都等于 n 时, 保存结果
		if lUsed == n && rUsed == n {
			res = append(res, ans)
			return
		}
		// 左括号没用完时总是可以添加左括号
		if lUsed < n {
			// 递归
			backtrack(ans+"(", lUsed+1, rUsed)
		}
		// 右括号只有在 左括号使用量大于右使用量且没用完时
		if rUsed < n && rUsed < lUsed {
			backtrack(ans+")", lUsed, rUsed+1)
		}
	}
	backtrack("", 0, 0)
	return res
}

func GenerateParenthesis(n int) []string {
	return generateParenthesis(n)
}
