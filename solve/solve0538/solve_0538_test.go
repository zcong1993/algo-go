package solve0538_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0538"
)

func TestConvertBST(t *testing.T) {
	input := tree.Deserialize("4,1,0,#,#,2,#,3,#,#,6,5,#,#,7,#,8,#,#")
	expected := "30,36,36,#,#,35,#,33,#,#,21,26,#,#,15,#,8,#,#"

	assert.Equal(t, expected, tree.Serialize(solve0538.ConvertBST(input)))
}
