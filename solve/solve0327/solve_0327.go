package solve0327

/**
@index 327
@title 区间和的个数
@difficulty 困难
@tags sort,binary-indexed-tree,segment-tree,binary-search,divide-and-conquer
@draft false
@link https://leetcode-cn.com/problems/count-of-range-sum/
@frontendId 327
*/
func countRangeSum(nums []int, lower int, upper int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i, val := range nums {
		preSum[i+1] = preSum[i] + val
	}
	ans := 0
	for i := 1; i < len(preSum); i++ {
		for j := 0; j < i; j++ {
			// sum [j, i-1]
			sum := preSum[i] - preSum[j]
			if sum >= lower && sum <= upper {
				ans++
			}
		}
	}
	return ans
}

func CountRangeSum(nums []int, lower int, upper int) int {
	return countRangeSum(nums, lower, upper)
}
