package solve0051

import "strings"

/**
@index 51
@title N 皇后
@difficulty 困难
@tags backtracking
@draft false
@link https://leetcode-cn.com/problems/n-queens/
@frontendId 51
*/

const (
	EMPTY = "."
	QUEEN = "Q"
)

func solveNQueens(n int) [][]string {
	board := makeBoard(n)
	res := make([][]string, 0)
	var backtrack func(m int)
	backtrack = func(m int) {
		// 到底了, 保存结果
		if m == n {
			res = append(res, toRes(board))
			return
		}
		// 尝试往当前行每个格子防止皇后
		for j := 0; j < n; j++ {
			// 不符合要求, 跳过
			if !isValid(board, m, j) {
				continue
			}
			// 放置皇后
			board[m][j] = QUEEN
			// 回溯
			backtrack(m + 1)
			// 撤销选择
			board[m][j] = EMPTY
		}
	}
	backtrack(0)
	return res
}

func isValid(board [][]string, m, n int) bool {
	l := len(board)
	for i := 0; i < l; i++ {
		if board[m][i] == QUEEN || board[i][n] == QUEEN {
			return false
		}
	}
	// 向下回溯, 只检查上方, 因为下方还没有元素
	// 右上 i-- j++
	for i, j := m-1, n+1; i >= 0 && j < l; i, j = i-1, j+1 {
		if board[i][j] == QUEEN {
			return false
		}
	}
	// 左上 i-- j--
	for i, j := m-1, n-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == QUEEN {
			return false
		}
	}
	return true
}

func makeBoard(n int) [][]string {
	res := make([][]string, n)
	for i := 0; i < n; i++ {
		line := make([]string, n)
		for j := 0; j < n; j++ {
			line[j] = EMPTY
		}
		res[i] = line
	}
	return res
}

func toRes(board [][]string) []string {
	res := make([]string, len(board))
	for i := 0; i < len(board); i++ {
		res[i] = strings.Join(board[i], "")
	}
	return res
}

func SolveNQueens(n int) [][]string {
	return solveNQueens(n)
}
