package solve0415_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0415"
)

func TestAddStrings(t *testing.T) {
	cases := []struct {
		num1     string
		num2     string
		expected string
	}{
		{
			"11",
			"123",
			"134",
		},
		{
			"456",
			"77",
			"533",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, solve0415.AddStrings(c.num1, c.num2))
	}
}
