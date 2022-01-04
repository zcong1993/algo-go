package solve0452

import "sort"

/**
 * @index 452
 * @title 用最少数量的箭引爆气球
 * @difficulty 中等
 * @tags greedy,array,sorting
 * @draft false
 * @link https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
 * @frontendId 452
 */

type interv [][]int

func (it interv) Len() int {
	return len(it)
}

func (it interv) Less(i, j int) bool {
	return it[i][1] < it[j][1]
}

func (it interv) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func findMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	inte := interv(points)
	sort.Sort(inte)
	res := 1
	end := inte[0][1]
	for _, in := range inte {
		if in[0] > end {
			res++
			end = in[1]
		}
	}
	return res
}
