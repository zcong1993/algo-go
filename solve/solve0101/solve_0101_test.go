package solve0101_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0101"
)

func TestIsSymmetric(t *testing.T) {
	root := tree.Deserialize("1,2,3,#,#,4,#,#,2,4,#,#,3,#,#")
	assert.True(t, solve0101.IsSymmetric(root))
}
