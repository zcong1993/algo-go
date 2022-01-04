package solve0069

import "math"

/**
 * @index 69
 * @title Sqrt(x)
 * @difficulty 简单
 * @tags math,binary-search
 * @draft false
 * @link https://leetcode-cn.com/problems/sqrtx/
 * @frontendId 69
 */

func MySqrt(x int) int {
	l, r := 0, x
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}

func MySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	c, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + c/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}
