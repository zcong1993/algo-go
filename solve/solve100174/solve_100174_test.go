package solve100174

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestSortedArrayToBST(t *testing.T) {
	root := SortedArrayToBST(helper.ArrFromJSON(`[-10,-3,0,5,9]`))
	assert.Equal(t, "0,-10,#,-3,#,#,5,#,9,#,#", tree.Serialize(root))
}
