package solve0903_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0903"
)

func TestRand10(t *testing.T) {
	for i := 0; i < 10000; i++ {
		r10 := solve0903.Rand10()
		assert.True(t, r10 >= 1 && r10 <= 10)
	}
}

func TestRand100(t *testing.T) {
	for i := 0; i < 10000; i++ {
		r10 := solve0903.Rand100()
		assert.True(t, r10 >= 1 && r10 <= 100)
	}
}
