package solve0538

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestConvertBST(t *testing.T) {
	root := tree.Deserialize("1,0,#,#,2,#,#")
	res := ConvertBST(root)
	assert.Equal(t, tree.Serialize(res), "3,3,#,#,2,#,#")
}
