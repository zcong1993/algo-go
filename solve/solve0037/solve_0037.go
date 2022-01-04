package solve0037

/**
 * @index 37
 * @title 解数独
 * @difficulty 困难
 * @tags array,backtracking,matrix
 * @draft false
 * @link https://leetcode-cn.com/problems/sudoku-solver/
 * @frontendId 37
 */

func isValid(board [][]byte, row, col int, val byte) bool {
	for i := 0; i < len(board); i++ {
		// 上下
		if board[i][col] == val {
			return false
		}
		// 左右
		if board[row][i] == val {
			return false
		}
	}
	rowS, colS := (row/3)*3, (col/3)*3

	for i := rowS; i < rowS+3; i++ {
		for j := colS; j < colS+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}

	return true
}

func SolveSudoku(board [][]byte) {
	l := len(board)
	var backtrack func(row, col int) bool
	backtrack = func(row, col int) bool {
		if col == l {
			return backtrack(row+1, 0)
		}

		if row == l {
			return true
		}

		if board[row][col] != '.' {
			return backtrack(row, col+1)
		}

		for i := 1; i <= 9; i++ {
			val := byte('0' + i)
			if !isValid(board, row, col, val) {
				continue
			}
			board[row][col] = val
			if backtrack(row, col+1) {
				return true
			}
			board[row][col] = '.'
		}
		return false
	}
	backtrack(0, 0)
}
