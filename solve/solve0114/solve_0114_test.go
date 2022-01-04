package solve0114_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0114"
)

func TestFlatten(t *testing.T) {
	input := tree.Deserialize("4,2,1,#,#,3,#,#,7,6,#,#,9,#,#")
	expected := "4,#,2,#,1,#,3,#,7,#,6,#,9,#,#"

	solve0114.Flatten(input)

	assert.Equal(t, expected, tree.Serialize(input))
}
