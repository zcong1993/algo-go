package solve0583_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0583"
)

func TestMinDistance(t *testing.T) {
	assert.Equal(t, 2, solve0583.MinDistance("sea", "eat"))
}
