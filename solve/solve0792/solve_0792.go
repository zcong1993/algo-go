package solve0792

/**
 * @index 792
 * @title 二分查找
 * @difficulty 简单
 * @tags array,binary-search
 * @draft false
 * @link https://leetcode-cn.com/problems/binary-search/
 * @frontendId 704
 */

// https://labuladong.gitee.io/algo/2/21/57/

func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	return -1
}
