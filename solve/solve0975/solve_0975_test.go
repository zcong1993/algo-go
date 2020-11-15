package solve0975

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestRangeSumBST(t *testing.T) {
	root := tree.Deserialize("10,5,3,#,#,7,#,#,15,#,18,#,#")
	assert.Equal(t, 32, RangeSumBST(root, 7, 15))
}
