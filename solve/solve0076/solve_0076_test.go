package solve0076_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve0076"
)

func TestMinWindow(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected string
	}{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"a",
			"a",
			"a",
		},
		{
			"a",
			"aa",
			"",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, solve0076.MinWindow(c.s, c.t))
	}
}
