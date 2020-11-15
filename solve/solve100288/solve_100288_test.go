package solve100288

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestMirrorTree(t *testing.T) {
	root := tree.Deserialize("4,2,1,#,#,3,#,#,7,6,#,#,9,#,#")
	root = MirrorTree(root)
	assert.Equal(t, "4,7,9,#,#,6,#,#,2,3,#,#,1,#,#", tree.Serialize(root))
}
