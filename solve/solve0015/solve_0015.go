package solve0015

import (
	"sort"
)

/**
 * @index 15
 * @title 三数之和
 * @difficulty 中等
 * @tags array,two-pointers,sorting
 * @draft false
 * @link https://leetcode-cn.com/problems/3sum/
 * @frontendId 15
 */

func TwoSum(nums []int, target int, start int) [][]int {
	var res [][]int
	l, h := start, len(nums)-1
	for l < h {
		left, right := nums[l], nums[h]
		sum := left + right
		if sum < target {
			l++
			for l < h && nums[l] == left {
				l++
			}
		} else if sum > target {
			h--
			for h > l && nums[h] == right {
				h--
			}
		} else if sum == target {
			res = append(res, []int{left, right})
			for l < h && nums[l] == left {
				l++
			}
			for h > l && nums[h] == right {
				h--
			}
		}
	}
	return res
}

func ThreeSum(nums []int) [][]int {
	var res [][]int

	if len(nums) < 3 {
		return res
	}

	target := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		twoRes := TwoSum(nums, target-cur, i+1)
		for _, r := range twoRes {
			res = append(res, append(r, cur))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
