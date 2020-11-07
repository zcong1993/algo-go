package solve0327

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func TestCountRangeSum(t *testing.T) {
	assert.Equal(t, 3, CountRangeSum(helper.ArrFromJSON("[-2,5,-1]"), -2, 2))
}
