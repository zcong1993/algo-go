package solve0037

/**
@index 37
@title 解数独
@difficulty 困难
@tags hash-table,backtracking
@draft false
@link https://leetcode-cn.com/problems/sudoku-solver/
@frontendId 37
*/

const (
	LEN   = 9
	EMPTY = '.'
)

func solveSudoku(board [][]byte) {
	var backtrack func(m, n int) bool
	backtrack = func(m, n int) bool {
		// 行结束时递归到下一行初
		if n == LEN {
			return backtrack(m+1, 0)
		}
		// 列结束时代表完成了
		if m == LEN {
			return true
		}
		for i := m; i < LEN; i++ {
			for j := n; j < LEN; j++ {
				// 已知元素不用处理, 直接递归到下个状态
				// 折行需要递归
				if board[i][j] != EMPTY {
					return backtrack(i, j+1)
				}
				// 尝试 1-9 所有可能
				for k := 0; k < LEN; k++ {
					char := byte(k + '1')
					// 检查能否将此值放在 i, j 上面, 不能直接向后
					if !isValid(board, i, j, char) {
						continue
					}
					// 选择
					board[i][j] = char
					// 回溯, 得到结果直接退出
					if backtrack(i, j+1) {
						return true
					}
					// 撤销
					board[i][j] = EMPTY
				}
				// 这个格子上所有元素都不合法, 无解
				return false
			}
		}
		// 所有格子都尝试了, 无解
		return false
	}
	backtrack(0, 0)
}

func isValid(board [][]byte, m, n int, val byte) bool {
	// 横竖
	for i := 0; i < LEN; i++ {
		if board[i][n] == val {
			return false
		}
		if board[m][i] == val {
			return false
		}
	}
	// 小九宫格
	mStart := m / 3 * 3
	nStart := n / 3 * 3
	for i := mStart; i < mStart+3; i++ {
		for j := nStart; j < nStart+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}

func SolveSudoku(board [][]byte) {
	solveSudoku(board)
}
