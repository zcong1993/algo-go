package solve0072

import . "github.com/zcong1993/algo-go/pkg/helper"

/**
 * @index 72
 * @title 编辑距离
 * @difficulty 困难
 * @tags string,dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/edit-distance/
 * @frontendId 72
 */

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// dp(i, j) -> minDistance(word1[0..i], word2[0..j])
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// 一个字符串为空时, 插入或者删除剩下的
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		} else {
			insert := dp(i, j-1)
			del := dp(i-1, j)
			modify := dp(i-1, j-1)
			return Min(insert, del, modify) + 1
		}
	}

	return dp(m-1, n-1)
}

func minDistanceMm(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	mm := make([][]int, m)
	for i := range mm {
		mm[i] = make([]int, n)
	}
	// dp(i, j) -> minDistance(word1[0..i], word2[0..j])
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// 一个字符串为空时, 插入或者删除剩下的
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if mm[i][j] != 0 {
			return mm[i][j]
		}
		if word1[i] == word2[j] {
			mm[i][j] = dp(i-1, j-1)
		} else {
			insert := dp(i, j-1)
			del := dp(i-1, j)
			modify := dp(i-1, j-1)
			mm[i][j] = Min(insert, del, modify) + 1
		}
		return mm[i][j]
	}

	return dp(m-1, n-1)
}

func minDistanceDp(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// dp[i-1][j-1] -> minDistance(word1[0..i], word2[0..j])
	// base case dp[0][j] = j dp[i][0] = i
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				insert := dp[i][j-1]
				del := dp[i-1][j]
				modify := dp[i-1][j-1]
				dp[i][j] = Min(insert, del, modify) + 1
			}
		}
	}

	return dp[m][n]
}

var (
	MinDistance   = minDistance
	MinDistanceMm = minDistanceMm
	MinDistanceDp = minDistanceDp
)
