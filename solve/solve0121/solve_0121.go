package solve0121

/**
 * @index 121
 * @title 买卖股票的最佳时机
 * @difficulty 简单
 * @tags array,dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
 * @frontendId 121
 */

func MaxProfit(prices []int) int {
	res := 0
	if len(prices) == 0 {
		return res
	}
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		cur := prices[i]
		r := cur - min
		if r > res {
			res = r
		}
		if cur < min {
			min = cur
		}
	}
	return res
}
