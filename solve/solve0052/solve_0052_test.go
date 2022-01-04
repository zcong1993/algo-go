package solve0052_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0052"
)

func TestTotalNQueens(t *testing.T) {
	assert.Equal(t, 2, solve0052.TotalNQueens(4))
	assert.Equal(t, 1, solve0052.TotalNQueens(1))
}
