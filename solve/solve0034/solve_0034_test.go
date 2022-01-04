package solve0034_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0034"
)

func TestSearchRange(t *testing.T) {
	input := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := []int{3, 4}
	assert.Equal(t, expected, solve0034.SearchRange(input, target))
}
