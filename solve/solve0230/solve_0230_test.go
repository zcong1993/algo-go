package solve0230_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0230"
)

func TestKthSmallest(t *testing.T) {
	input := tree.Deserialize("3,1,#,2,#,#,4,#,#")

	for i := 1; i <= 4; i++ {
		assert.Equal(t, i, solve0230.KthSmallest(input, i))
	}
}
