package solve0712

import . "github.com/zcong1993/algo-go/pkg/helper"

/**
 * @index 712
 * @title 两个字符串的最小ASCII删除和
 * @difficulty 中等
 * @tags dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings/
 * @frontendId 712
 */

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	mm := make([][]int, m)
	for i := range mm {
		mm[i] = make([]int, n)
	}

	getRestSum := func(s string, i int) int {
		ans := 0
		for i < len(s) {
			ans += int(s[i])
			i++
		}
		return ans
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == m {
			return getRestSum(s2, j)
		}
		if j == n {
			return getRestSum(s1, i)
		}
		if mm[i][j] != 0 {
			return mm[i][j]
		}
		if s1[i] == s2[j] {
			mm[i][j] = dp(i+1, j+1)
		} else {
			mm[i][j] = Min2(int(s1[i])+dp(i+1, j), int(s2[j])+dp(i, j+1))
		}
		return mm[i][j]
	}

	return dp(0, 0)
}

func minimumDeleteSumDp(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 一个字符串为空时, 值为剩下字符串 code 之和
	for i := m - 1; i >= 0; i-- {
		dp[i][n] = dp[i+1][n] + int(s1[i])
	}
	for j := n - 1; j >= 0; j-- {
		dp[m][j] = dp[m][j+1] + int(s2[j])
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = Min2(dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j]))
			}
		}
	}
	return dp[0][0]
}

var (
	MinimumDeleteSum   = minimumDeleteSum
	MinimumDeleteSumDp = minimumDeleteSumDp
)
