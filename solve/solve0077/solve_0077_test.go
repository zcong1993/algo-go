package solve0077

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func TestCombine(t *testing.T) {
	res := Combine(4, 2)
	assert.Equal(t, helper.GridFromJSON(`[[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]`), res)
}
