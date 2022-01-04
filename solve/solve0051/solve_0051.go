package solve0051

/**
 * @index 51
 * @title N 皇后
 * @difficulty 困难
 * @tags array,backtracking
 * @draft false
 * @link https://leetcode-cn.com/problems/n-queens/
 * @frontendId 51
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

func copyBoard(board [][]byte) []string {
	l := len(board)
	res := make([]string, l)
	for i := 0; i < l; i++ {
		line := string(board[i])
		res[i] = line
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

func SolveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}

	var res [][]string

	board := makeBoard(n)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == len(board) {
			res = append(res, copyBoard(board))
			return
		}
		for i := 0; i < len(board); i++ {
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
