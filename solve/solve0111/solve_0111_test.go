package solve0111_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0111"
)

func TestMinDepth(t *testing.T) {
	root1 := tree.Deserialize("3,9,#,#,20,15,#,#,7,#,#")
	assert.Equal(t, 2, solve0111.MinDepth(root1))

	root2 := tree.Deserialize("3,#,2,#,#")
	assert.Equal(t, 2, solve0111.MinDepth(root2))

	root3 := tree.Deserialize("3,2,#,#,#")
	assert.Equal(t, 2, solve0111.MinDepth(root3))

	assert.Equal(t, 0, solve0111.MinDepth(nil))
}
