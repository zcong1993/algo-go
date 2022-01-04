package solve0199_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0199"
)

func TestRightSideView(t *testing.T) {
	root := tree.Deserialize("1,2,#,5,#,#,3,#,4,#,#")
	expected := []int{1, 3, 4}

	assert.Equal(t, expected, solve0199.RightSideView(root))
}
