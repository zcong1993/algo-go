package solve0654_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0654"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	input := helper.ArrFromJSON(`[3,2,1,6,0,5]`)
	expected := "6,3,#,2,#,1,#,#,5,0,#,#,#"

	tr := solve0654.ConstructMaximumBinaryTree(input)

	assert.Equal(t, expected, tree.Serialize(tr))
}
