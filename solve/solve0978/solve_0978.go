package solve0978

/*
*
@index 978
@title 有效的山脉数组
@difficulty 简单
@tags array
@draft false
@link https://leetcode-cn.com/problems/valid-mountain-array/
@frontendId 941.
*/
func ValidMountainArray(A []int) bool {
	if len(A) == 0 {
		return false
	}
	// 递增 true 还是递减 false
	flag := true
	// 前一个值
	prev := A[0]
	// 有没有递增序列, 只有递减应该是 false
	hasIncr := false
	for i := 1; i < len(A); i++ {
		// 递减序列时, 出现递增, return false
		if !flag && A[i] >= prev {
			return false
		}

		// 递增序列时
		if flag {
			// 发现第一个递减, 改变 flag
			if A[i] < prev {
				flag = false
			} else if A[i] > prev { // 标记有递增序列
				hasIncr = true
			}
		}
		// 更新前一个值
		prev = A[i]
	}
	// 当前是递减序列而且有过递增序列
	return !flag && hasIncr
}

func ValidMountainArray2(A []int) bool {
	if len(A) == 0 {
		return false
	}
	i := 0
	// 做到右遍历递增序列, 得到顶点 index
	for ; i < len(A)-1 && A[i] < A[i+1]; i++ {
	}
	// index 不能是 0 或者 末尾, 因为必须同时有递增和递减
	if i == 0 || i == len(A)-1 {
		return false
	}
	// 继续遍历递减序列
	for ; i < len(A)-1 && A[i] > A[i+1]; i++ {
	}
	// 最终 index 等于末尾时, return true
	return i == len(A)-1
}
