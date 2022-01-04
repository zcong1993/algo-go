package solve0034

/**
 * @index 34
 * @title 在排序数组中查找元素的第一个和最后一个位置
 * @difficulty 中等
 * @tags array,binary-search
 * @draft false
 * @link https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
 * @frontendId 34
 */

// https://labuladong.gitee.io/algo/2/21/57

func searchLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			r = mid - 1
		}
	}

	if l >= len(nums) || nums[l] != target {
		return -1
	}
	return l
}

func searchRight(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			l = mid + 1
		}
	}

	if r < 0 || nums[r] != target {
		return -1
	}
	return r
}

func SearchRange(nums []int, target int) []int {
	return []int{searchLeft(nums, target), searchRight(nums, target)}
}
