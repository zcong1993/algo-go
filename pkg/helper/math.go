package helper

// MaxInt is max safe int
const MaxInt = int(^uint(0) >> 1)

// MinInt is min safe int
const MinInt = -MaxInt - 1

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
