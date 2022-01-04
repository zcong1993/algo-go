package solve0215

import (
	"math/rand"

	. "github.com/zcong1993/algo-go/pkg/pq"
)

/**
 * @index 215
 * @title 数组中的第K个最大元素
 * @difficulty 中等
 * @tags array,divide-and-conquer,quickselect,sorting,heap-priority-queue
 * @draft false
 * @link https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
 * @frontendId 215
 */

func FindKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}

	p := NewPQ(len(nums), MIN)
	for _, num := range nums {
		p.Add(num)
		if p.Size() > k {
			p.DelTop()
		}
	}

	return p.Peek()
}

// quick select
// 1. 利用 quick sort partition 能够确保 [l, p-1] < p < [p+1, h]
// 2. 倒数第 k 大的元素为升序排序后数组索引为 len - k (k') 元素
// 3. 二分搜索, p < k' 时搜索右半部分, 反之搜索左边
// 4. partition 注意随机选择元素

func FindKthLargest2(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}
	kk := len(nums) - k
	i, j := 0, len(nums)-1
	for i <= j {
		p := partition(nums, i, j)
		if p < kk {
			i = p + 1
		} else if p > kk {
			j = p - 1
		} else {
			return nums[p]
		}
	}
	return -1
}

func swap(nums []int, l, r int) {
	if l == r {
		return
	}
	nums[l], nums[r] = nums[r], nums[l]
}

func partition(nums []int, l, r int) int {
	if l == r {
		return l
	}
	pIdx := l + rand.Intn(r-l)
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
