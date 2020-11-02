package solve0105

import (
	"testing"

	"github.com/zcong1993/algo-go/pkg/helper"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestBuildTree(t *testing.T) {
	expected := "1,2,4,#,#,5,#,#,3,6,#,#,7,#,#"
	root := BuildTree(helper.ArrFromJSON("[1,2,4,5,3,6,7]"), helper.ArrFromJSON("[4,2,5,1,6,3,7]"))
	assert.Equal(t, expected, tree.Serialize(root))
}
