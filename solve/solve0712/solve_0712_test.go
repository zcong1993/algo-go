package solve0712_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0712"
)

func TestMinimumDeleteSum(t *testing.T) {
	assert.Equal(t, 231, solve0712.MinimumDeleteSum("sea", "eat"))
	assert.Equal(t, 231, solve0712.MinimumDeleteSumDp("sea", "eat"))
}
