package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 5, Max(1, 2, 3, 4, 5, 0, -10))
}

func TestMin(t *testing.T) {
	assert.Equal(t, -10, Min(1, 2, 3, 4, 5, 0, -10))
}

func TestMin2(t *testing.T) {
	assert.Equal(t, -1, Min2(-1, 0))
}

func TestMax2(t *testing.T) {
	assert.Equal(t, 0, Max2(-1, 0))
}
