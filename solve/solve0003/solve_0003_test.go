package solve0003_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0003"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, solve0003.LengthOfLongestSubstring(c.input))
	}
}
