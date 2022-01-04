package solve0322_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0322"
)

func testCoinChange(t *testing.T, fn func(coins []int, amount int) int) {
	cases := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, fn(c.coins, c.amount))
	}
}

func TestCoinChange(t *testing.T) {
	testCoinChange(t, solve0322.CoinChange)
}

func TestCoinChange2(t *testing.T) {
	testCoinChange(t, solve0322.CoinChange2)
}

func TestCoinChange3(t *testing.T) {
	testCoinChange(t, solve0322.CoinChange3)
}
