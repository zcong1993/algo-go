package solve0435_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0435"
)

func TestEraseOverlapIntervals(t *testing.T) {
	inputs := helper.GridFromJSON(`[ [1,2], [2,3], [3,4], [1,3] ]`)
	expected := 1
	assert.Equal(t, expected, solve0435.EraseOverlapIntervals(inputs))
}
