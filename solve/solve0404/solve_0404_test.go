package solve0404_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0404"
)

func TestSumOfLeftLeaves(t *testing.T) {
	input := tree.Deserialize("3,9,#,#,20,15,#,#,7,#,#")
	expected := 24
	assert.Equal(t, expected, solve0404.SumOfLeftLeaves(input))

	input = tree.Deserialize("3,#,#")
	expected = 0
	assert.Equal(t, expected, solve0404.SumOfLeftLeaves(input))
}
