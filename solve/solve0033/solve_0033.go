package solve0033

/**
 * @index 33
 * @title 搜索旋转排序数组
 * @difficulty 中等
 * @tags array,binary-search
 * @draft false
 * @link https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
 * @frontendId 33
 */

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		lVal, rVal, mVal := nums[l], nums[r], nums[m]

		// 将搜索到的结果提前
		if mVal == target {
			return m
		}

		// 数组有序, 没有发生旋转, 标准二分
		if mVal >= lVal && mVal <= rVal {
			if target < mVal {
				r = m - 1
			} else if target > mVal {
				l = m + 1
			}
		} else { // 发生旋转
			// 优化: 发生旋转后左右边界元素应该是连续的且 lVal > rVal
			// 处于他们之间的元素肯定不存在
			if target > rVal && target < lVal {
				return -1
			}

			// 等于边界的情况直接返回
			if target == lVal {
				return l
			}
			if target == rVal {
				return r
			}

			// 旋转发生在左边
			if lVal > mVal {
				// 右边有序, 搜索右边
				if target > mVal && target < rVal {
					l = m + 1
				} else {
					r = m - 1
				}
			} else {
				if target > lVal && target < mVal {
					r = m - 1
				} else {
					l = m + 1
				}
			}
		}
	}

	return -1
}
