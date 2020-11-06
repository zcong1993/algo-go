package solve1458

import "sort"

/**
@index 1458
@title 根据数字二进制下 1 的数目排序
@difficulty 简单
@tags sort,bit-manipulation
@draft false
@link https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/
@frontendId 1356
*/

type ByBits []int

func (a ByBits) Len() int      { return len(a) }
func (a ByBits) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByBits) Less(i, j int) bool {
	if countBits(a[i]) == countBits(a[j]) {
		return a[i] < a[j]
	}
	return countBits(a[i]) < countBits(a[j])
}

var mm map[int]int

func sortByBits(arr []int) []int {
	mm = make(map[int]int, len(arr))
	var bb ByBits
	bb = arr
	sort.Sort(bb)
	return bb
}

func countBits(x int) int {
	origin := x
	if n, ok := mm[x]; ok {
		return n
	}
	ans := 0
	for x > 0 {
		x -= x & -x
		ans++
	}
	mm[origin] = ans
	return ans
}

func SortByBits(arr []int) []int {
	return sortByBits(arr)
}
