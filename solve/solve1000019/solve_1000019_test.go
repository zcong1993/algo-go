package solve1000019_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve1000019"
)

func TestConvertBiNode(t *testing.T) {
	input := tree.Deserialize("4,2,1,0,#,#,#,3,#,#,5,#,6,#,#")
	expected := "0,#,1,#,2,#,3,#,4,#,5,#,6,#,#"

	res := solve1000019.ConvertBiNode(input)

	assert.Equal(t, expected, tree.Serialize(res))
}
