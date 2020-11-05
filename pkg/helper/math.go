package helper

import "math"

// MaxInt is max safe int
const MaxInt int = math.MaxInt64

// MinInt is min safe int
const MinInt int = math.MinInt64

// Max find the biggest num of nums
func Max(nums ...int) int {
	max := MinInt
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// Min find minist num of nums
func Min(nums ...int) int {
	min := MaxInt
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// Min2 return min num of 2 nums
func Min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max2 return max num of 2 nums
func Max2(a, b int) int {
	if a < b {
		return b
	}
	return a
}
