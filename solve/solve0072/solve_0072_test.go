package solve0072_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0072"
)

func testMinDistance(t *testing.T, f func(word1 string, word2 string) int) {
	assert.Equal(t, 3, f("horse", "ros"))
	assert.Equal(t, 5, f("intention", "execution"))
}

func TestMinDistance(t *testing.T) {
	testMinDistance(t, solve0072.MinDistance)
	testMinDistance(t, solve0072.MinDistanceMm)
	testMinDistance(t, solve0072.MinDistanceDp)
}
