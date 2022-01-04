package solve0111_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0111"
)

func TestMinDepth(t *testing.T) {
	input := tree.Deserialize("2,#,3,#,4,#,#")
	expected := 3
	assert.Equal(t, expected, solve0111.MinDepth(input))

	input = tree.Deserialize("3,9,#,#,20,15,#,#,7,#,#")
	expected = 2
	assert.Equal(t, expected, solve0111.MinDepth(input))
}
