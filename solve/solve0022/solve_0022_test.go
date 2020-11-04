package solve0022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	res := GenerateParenthesis(3)
	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, res)
}
