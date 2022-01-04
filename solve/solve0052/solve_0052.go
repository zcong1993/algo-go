package solve0052

/**
 * @index 52
 * @title N皇后 II
 * @difficulty 困难
 * @tags backtracking
 * @draft false
 * @link https://leetcode-cn.com/problems/n-queens-ii/
 * @frontendId 52
 */

const (
	EMPTY = '.'
	QUEEN = 'Q'
)

func makeBoard(n int) [][]byte {
	res := make([][]byte, n)
	line := make([]byte, n)
	for i := range line {
		line[i] = EMPTY
	}
	res[0] = line
	for i := 1; i < n; i++ {
		tmp := make([]byte, n)
		copy(tmp, line)
		res[i] = tmp
	}
	return res
}

func isValid(board [][]byte, row, col int) bool {
	l := len(board)
	// 上
	for i := 0; i < row; i++ {
		if board[i][col] == QUEEN {
			return false
		}
	}
	// 右上
	for i, j := row-1, col+1; i >= 0 && j < l; i, j = i-1, j+1 {
		if board[i][j] == QUEEN {
			return false
		}
	}
	// 左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == QUEEN {
			return false
		}
	}

	return true
}

func TotalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	board := makeBoard(n)
	res := 0
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			res++
			return
		}
		for i := 0; i < n; i++ {
			if !isValid(board, row, i) {
				continue
			}
			board[row][i] = QUEEN
			backtrack(row + 1)
			board[row][i] = EMPTY
		}
	}
	backtrack(0)
	return res
}
