package solve0747_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0747"
)

func TestMinCostClimbingStairs(t *testing.T) {
	assert.Equal(t, 15, solve0747.MinCostClimbingStairs(helper.ArrFromJSON(`[10, 15, 20]`)))
	assert.Equal(t, 6, solve0747.MinCostClimbingStairs(helper.ArrFromJSON(`[1, 100, 1, 1, 1, 100, 1, 1, 100, 1]`)))
}

func TestMinCostClimbingStairs2(t *testing.T) {
	assert.Equal(t, 15, solve0747.MinCostClimbingStairs2(helper.ArrFromJSON(`[10, 15, 20]`)))
	assert.Equal(t, 6, solve0747.MinCostClimbingStairs2(helper.ArrFromJSON(`[1, 100, 1, 1, 1, 100, 1, 1, 100, 1]`)))
}
