package solve0349

/**
@index 349
@title 两个数组的交集
@difficulty 简单
@tags sort,hash-table,two-pointers,binary-search
@draft false
@link https://leetcode-cn.com/problems/intersection-of-two-arrays/
*/
func Intersection(nums1 []int, nums2 []int) []int {
	mm := make(map[int]struct{}, len(nums1))
	for _, n := range nums1 {
		mm[n] = struct{}{}
	}
	resMM := make(map[int]struct{}, 0)
	for _, n := range nums2 {
		if _, ok := mm[n]; ok {
			resMM[n] = struct{}{}
		}
	}
	res, i := make([]int, len(resMM)), 0
	for k := range resMM {
		res[i] = k
		i++
	}
	return res
}
