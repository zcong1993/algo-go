package solve1000019

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestConvertBiNode(t *testing.T) {
	root := tree.Deserialize("1,2,#,#,3,#,#")
	res := ConvertBiNode(root)
	assert.Equal(t, "2,#,1,#,3,#,#", tree.Serialize(res))
}
