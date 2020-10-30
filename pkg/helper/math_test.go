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
