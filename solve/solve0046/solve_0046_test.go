package solve0046

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func TestPermute(t *testing.T) {
	assert.Equal(t, helper.GridFromJSON(`[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]`), Permute(helper.ArrFromJSON(`[1,2,3]`)))
}
