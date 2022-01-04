package solve1013

/**
 * @index 1013
 * @title 斐波那契数
 * @difficulty 简单
 * @tags recursion,memoization,math,dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/fibonacci-number/
 * @frontendId 509
 */

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Fib2(n int) int {
	return fib2(n, make(map[int]int))
}

func fib2(n int, store map[int]int) int {
	if res, ok := store[n]; ok {
		return res
	}
	var res int
	if n < 2 {
		res = n
	} else {
		res = fib2(n-1, store) + fib2(n-2, store)
	}
	store[n] = res
	return res
}

func Fib3(n int) int {
	if n < 2 {
		return n
	}
	n2, n1 := 0, 1
	for i := 2; i <= n; i++ {
		next := n2 + n1
		n2, n1 = n1, next
	}
	return n1
}

func Fib4(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
