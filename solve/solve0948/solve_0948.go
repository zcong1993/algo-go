package solve0948

import "math/rand"

/**
 * @index 948
 * @title 排序数组
 * @difficulty 中等
 * @tags array,divide-and-conquer,bucket-sort,counting-sort,radix-sort,sorting,heap-priority-queue,merge-sort
 * @draft false
 * @link https://leetcode-cn.com/problems/sort-an-array/
 * @frontendId 912
 */

func SortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func swap(nums []int, l, r int) {
	nums[l], nums[r] = nums[r], nums[l]
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
	if l == r {
		return l
	}
	pIdx := rand.Intn(r-l) + l
	p := nums[pIdx]
	swap(nums, pIdx, r)
	idx := l - 1
	for i := l; i < r; i++ {
		if nums[i] < p {
			idx++
			swap(nums, idx, i)
		}
	}
	idx++
	swap(nums, idx, r)
	return idx
}
