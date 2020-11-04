package solve0057

import . "github.com/zcong1993/algo-go/pkg/helper"

/**
@index 57
@title 插入区间
@difficulty 困难
@tags sort,array
@draft false
@link https://leetcode-cn.com/problems/insert-interval/
@frontendId 57
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	newL, newR := newInterval[0], newInterval[1]
	merged := false
	res := make([][]int, 0)
	for _, interval := range intervals {
		l, r := interval[0], interval[1]
		// 插入区间在当前区间右边, 直接保存当前区间
		if r < newL {
			res = append(res, interval)
			//	插入区间在当前区间左边
		} else if l > newR {
			// 如果没有合并, 将新区间加入结果
			if !merged {
				res = append(res, []int{newL, newR})
				merged = true
			}
			// 将当前区间加入结果
			res = append(res, interval)
			// 区间有重合, 合并更新新区间
		} else {
			newL = Min2(l, newL)
			newR = Max2(r, newR)
		}
	}
	// 如果还没有合并, 将新区间加入结果
	if !merged {
		res = append(res, []int{newL, newR})
	}
	return res
}

func Insert(intervals [][]int, newInterval []int) [][]int {
	return insert(intervals, newInterval)
}
