package solve1250

import . "github.com/zcong1993/algo-go/pkg/helper"

/**
 * @index 1250
 * @title 最长公共子序列
 * @difficulty 中等
 * @tags dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/longest-common-subsequence/
 * @frontendId 1143
 */

// 递归方法
func longestCommonSubsequence(text1 string, text2 string) int {
	// func(i, j int) 表示 text1[i:] text2[j:] 最长公共子序列
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// 一个字符串长度为 0, 所以返回 0
		if i == len(text1) || j == len(text2) {
			return 0
		}
		// 当前字符相等 i,j 都向右移动
		if text1[i] == text2[j] {
			return 1 + dp(i+1, j+1)
		} else {
			// 三种情况取最大 1. i+1, j 2. i, j+1 3. i+1, j+1
			// 第三种被前两种包含, 所以可以忽略
			return Max2(dp(i+1, j), dp(i, j+1))
		}
	}
	return dp(0, 0)
}

// 备忘录
func longestCommonSubsequence2(text1 string, text2 string) int {
	mm := make([][]int, len(text1))
	for i := range mm {
		mm[i] = make([]int, len(text2))
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}
		if val := mm[i][j]; val != 0 {
			return val
		}
		var res int
		if text1[i] == text2[j] {
			res = 1 + dp(i+1, j+1)
		} else {
			res = Max2(dp(i+1, j), dp(i, j+1))
		}
		mm[i][j] = res
		return res
	}
	return dp(0, 0)
}

// dp
func longestCommonSubsequence3(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	// 防止越界, index+1
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = Max2(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

var (
	LongestCommonSubsequence  = longestCommonSubsequence
	LongestCommonSubsequence2 = longestCommonSubsequence2
	LongestCommonSubsequence3 = longestCommonSubsequence3
)
