package solve0053_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0053"
)

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, 6, solve0053.MaxSubArray(helper.ArrFromJSON(`[-2,1,-3,4,-1,2,1,-5,4]`)))
}
