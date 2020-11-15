package solve100319

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestMaxDepth(t *testing.T) {
	root := tree.Deserialize("3,9,#,#,20,15,#,#,7,#,#")
	assert.Equal(t, 3, MaxDepth(root))
}
