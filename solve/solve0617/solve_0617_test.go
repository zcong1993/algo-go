package solve0617_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0617"
)

func TestMergeTrees(t *testing.T) {
	input1 := tree.Deserialize("1,3,5,#,#,#,2,#,#")
	input2 := tree.Deserialize("2,1,#,4,#,#,3,#,7,#,#")
	expected := "3,4,5,#,#,4,#,#,5,#,7,#,#"

	assert.Equal(t, expected, tree.Serialize(solve0617.MergeTrees(input1, input2)))
}
