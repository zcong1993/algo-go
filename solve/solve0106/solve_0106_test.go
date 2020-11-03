package solve0106

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestBuildTree(t *testing.T) {
	root := BuildTree(helper.ArrFromJSON("[9,3,15,20,7]"), helper.ArrFromJSON("[9,15,7,20,3]"))
	assert.Equal(t, "3,9,#,#,20,15,#,#,7,#,#", tree.Serialize(root))
}
