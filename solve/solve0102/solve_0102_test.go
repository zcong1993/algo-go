package solve0102

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestLevelOrder(t *testing.T) {
	root := tree.Deserialize("3,9,#,#,20,15,#,#,7,#,#")
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, LevelOrder(root))
}
