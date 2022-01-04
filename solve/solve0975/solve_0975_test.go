package solve0975_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0975"
)

func TestRangeSumBST(t *testing.T) {
	input := tree.Deserialize("10,5,3,#,#,7,#,#,15,#,18,#,#")
	low, high := 7, 15
	expected := 32
	assert.Equal(t, expected, solve0975.RangeSumBST(input, low, high))
}

func TestRangeSumBST2(t *testing.T) {
	input := tree.Deserialize("10,5,3,#,#,7,#,#,15,#,18,#,#")
	low, high := 7, 15
	expected := 32
	assert.Equal(t, expected, solve0975.RangeSumBST2(input, low, high))
}
