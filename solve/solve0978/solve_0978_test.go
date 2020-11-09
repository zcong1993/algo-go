package solve0978

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func TestValidMountainArray(t *testing.T) {
	assert.True(t, ValidMountainArray(helper.ArrFromJSON("[0,3,2,1]")))
	assert.False(t, ValidMountainArray(helper.ArrFromJSON("[3,5,5]")))
	assert.False(t, ValidMountainArray(helper.ArrFromJSON("[2,1]")))
	assert.False(t, ValidMountainArray([]int{}))
	assert.False(t, ValidMountainArray(helper.ArrFromJSON("[3,5,4,6]")))
}

func TestValidMountainArray2(t *testing.T) {
	assert.True(t, ValidMountainArray2(helper.ArrFromJSON("[0,3,2,1]")))
	assert.False(t, ValidMountainArray2(helper.ArrFromJSON("[3,5,5]")))
	assert.False(t, ValidMountainArray2(helper.ArrFromJSON("[2,1]")))
	assert.False(t, ValidMountainArray2([]int{}))
	assert.False(t, ValidMountainArray2(helper.ArrFromJSON("[3,5,4,6]")))
}
