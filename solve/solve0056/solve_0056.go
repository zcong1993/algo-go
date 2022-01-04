package solve0056

import (
	"sort"

	. "github.com/zcong1993/algo-go/pkg/helper"
)

/**
 * @index 56
 * @title 合并区间
 * @difficulty 中等
 * @tags array,sorting
 * @draft false
 * @link https://leetcode-cn.com/problems/merge-intervals/
 * @frontendId 56
 */

type inter [][]int

func (in inter) Len() int {
	return len(in)
}

func (in inter) Less(i, j int) bool {
	return in[i][0] < in[j][0]
}

func (in inter) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	in := inter(intervals)
	sort.Sort(in)
	res := [][]int{in[0]}
	for i := 1; i < len(in); i++ {
		last := res[len(res)-1]
		cur := in[i]
		if cur[0] <= last[1] {
			last[1] = Max2(last[1], cur[1])
		} else {
			res = append(res, cur)
		}
	}
	return res
}
