package solve0322

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func TestCoinChange(t *testing.T) {
	assert.Equal(t, 3, CoinChange(helper.ArrFromJSON("[1, 2, 5]"), 11))
	assert.Equal(t, -1, coinChange([]int{2}, 3))
	assert.Equal(t, 0, coinChange([]int{1}, 0))
	assert.Equal(t, 1, coinChange([]int{1}, 1))
}
