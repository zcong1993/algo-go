package solve0572

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestIsSubtree(t *testing.T) {
	s1 := tree.Deserialize("3,4,1,#,#,2,#,#,5,#,#")
	t1 := tree.Deserialize("4,1,#,#,2,#,#")
	assert.True(t, IsSubtree(s1, t1))

	s2 := tree.Deserialize("3,4,1,#,#,2,0,#,#,#,5,#,#")
	t2 := tree.Deserialize("4,1,#,#,2,#,#")
	assert.False(t, IsSubtree(s2, t2))

	assert.True(t, IsSubtree(nil, nil))
}
