package solve0042

import "math"

/**
 * @index 42
 * @title 接雨水
 * @difficulty 困难
 * @tags stack,array,two-pointers,dynamic-programming,monotonic-stack
 * @draft false
 * @link https://leetcode-cn.com/problems/trapping-rain-water/
 * @frontendId 42
 */

func Trap(height []int) int {
	res := 0
	lMax, rMax := math.MinInt32, math.MinInt32
	l, r := 0, len(height)-1
	for l < r {
		lVal, rVal := height[l], height[r]

		if lVal > lMax {
			lMax = lVal
		}
		if rVal > rMax {
			rMax = rVal
		}

		if lMax <= rMax {
			res += lMax - lVal
			l++
		} else {
			res += rMax - rVal
			r--
		}
	}
	return res
}
