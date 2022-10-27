package solve0449_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0449"
)

var (
	testTreeStr = "4,2,1,#,#,3,#,#,7,6,#,#,9,#,#"
	testTree    = tree.Deserialize(testTreeStr)
)

var c = solve0449.Constructor()

func TestCodec_Serialize(t *testing.T) {
	assert.Equal(t, testTreeStr, c.Serialize(testTree))
}

func TestCodec_Deserialize(t *testing.T) {
	assert.True(t, tree.IsSameTree(testTree, c.Deserialize(testTreeStr)))
}
