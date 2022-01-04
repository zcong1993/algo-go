package solve0435

import "sort"

/**
 * @index 435
 * @title 无重叠区间
 * @difficulty 中等
 * @tags greedy,array,dynamic-programming,sorting
 * @draft false
 * @link https://leetcode-cn.com/problems/non-overlapping-intervals/
 * @frontendId 435
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

func intervalSchedule(intervals [][]int) int {
	if len(intervals) < 2 {
		return len(intervals)
	}

	inte := interv(intervals)
	sort.Sort(inte)
	res := 1
	end := inte[0][1]
	for _, in := range inte {
		if in[0] >= end {
			res++
			end = in[1]
		}
	}
	return res
}

func EraseOverlapIntervals(intervals [][]int) int {
	return len(intervals) - intervalSchedule(intervals)
}
