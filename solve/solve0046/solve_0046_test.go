package solve0046_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0046"
)

func TestPermute(t *testing.T) {
	inputs := []int{1, 2, 3}
	expected := helper.GridFromJSON(`[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]`)
	assert.Equal(t, expected, solve0046.Permute(inputs))
}
