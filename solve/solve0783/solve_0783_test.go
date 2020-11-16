package solve0783_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0783"
)

func TestSearchBST(t *testing.T) {
	root := tree.Deserialize("4,2,1,#,#,3,#,#,7,#,#")
	assert.Equal(t, "2,1,#,#,3,#,#", tree.Serialize(solve0783.SearchBST(root, 2)))
	assert.Nil(t, solve0783.SearchBST(root, 5))
	assert.Equal(t, "7,#,#", tree.Serialize(solve0783.SearchBST(root, 7)))
}
