package solve0799_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0799"
)

func TestMinDiffInBST(t *testing.T) {
	input := tree.Deserialize("4,2,1,#,#,3,#,#,6,#,#")
	expected := 1
	assert.Equal(t, expected, solve0799.MinDiffInBST(input))

	input = tree.Deserialize("1,0,#,#,48,12,#,#,49,#,#")
	expected = 1
	assert.Equal(t, expected, solve0799.MinDiffInBST(input))
}
