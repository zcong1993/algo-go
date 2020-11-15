package solve0617

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestMergeTrees(t *testing.T) {
	t1 := tree.Deserialize("1,3,5,#,#,#,2,#,#")
	t2 := tree.Deserialize("2,1,#,4,#,#,3,#,7,#,#")
	assert.Equal(t, "3,4,5,#,#,4,#,#,5,#,7,#,#", tree.Serialize(mergeTrees(t1, t2)))
}
