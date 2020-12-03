package solve0583

import . "github.com/zcong1993/algo-go/solve/solve1250"

/**
 * @index 583
 * @title 两个字符串的删除操作
 * @difficulty 中等
 * @tags string
 * @draft false
 * @link https://leetcode-cn.com/problems/delete-operation-for-two-strings/
 * @frontendId 583
 */

func minDistance(word1 string, word2 string) int {
	n := LongestCommonSubsequence3(word1, word2)
	return len(word1) - n + len(word2) - n
}

var MinDistance = minDistance
