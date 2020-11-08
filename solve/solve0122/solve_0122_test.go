package solve0122

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 7, MaxProfit(helper.ArrFromJSON("[7,1,5,3,6,4]")))
	assert.Equal(t, 0, MaxProfit([]int{}))
}

func TestMaxProfitDp(t *testing.T) {
	assert.Equal(t, 7, MaxProfitDp(helper.ArrFromJSON("[7,1,5,3,6,4]")))
	assert.Equal(t, 0, MaxProfitDp([]int{}))
}

func TestMaxProfitDpOptim(t *testing.T) {
	assert.Equal(t, 7, MaxProfitDpOptim(helper.ArrFromJSON("[7,1,5,3,6,4]")))
	assert.Equal(t, 0, MaxProfitDpOptim([]int{}))
}
