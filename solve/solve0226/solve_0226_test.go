package solve0226_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0226"
)

func TestInvertTree(t *testing.T) {
	input := tree.Deserialize("4,2,1,#,#,3,#,#,7,6,#,#,9,#,#")
	expected := "4,7,9,#,#,6,#,#,2,3,#,#,1,#,#"

	assert.Equal(t, expected, tree.Serialize(solve0226.InvertTree(input)))
}
