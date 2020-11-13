package solve0198

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func testRob(t *testing.T, f func(nums []int) int) {
	assert.Equal(t, 4, f(helper.ArrFromJSON(`[1,2,3,1]`)))
}

func TestRob(t *testing.T) {
	testRob(t, Rob)
}

func TestRobDp(t *testing.T) {
	testRob(t, RobDp)
}

func TestRobDpOptim(t *testing.T) {
	testRob(t, RobDpOptim)
}
