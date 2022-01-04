package solve0124_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0124"
)

func TestMaxPathSum(t *testing.T) {
	inputs := tree.Deserialize("1,2,#,#,3,#,#")
	expected := 6
	assert.Equal(t, expected, solve0124.MaxPathSum(inputs))

	inputs = tree.Deserialize("-10,9,#,#,20,15,#,#,7,#,#")
	expected = 42
	assert.Equal(t, expected, solve0124.MaxPathSum(inputs))
}
