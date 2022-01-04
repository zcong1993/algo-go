package solve0069_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0069"
)

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 2, solve0069.MySqrt(4))
	assert.Equal(t, 2, solve0069.MySqrt(8))
}

func TestMySqrt2(t *testing.T) {
	assert.Equal(t, 2, solve0069.MySqrt2(4))
	assert.Equal(t, 2, solve0069.MySqrt2(8))
}
