package solve100304

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, 6, MaxSubArray(helper.ArrFromJSON(`[-2,1,-3,4,-1,2,1,-5,4]`)))
}
