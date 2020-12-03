package solve1250_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve1250"
)

func testLongestCommonSubsequence(t *testing.T, f func(text1 string, text2 string) int) {
	assert.Equal(t, 3, f("abcde", "ace"))
	assert.Equal(t, 3, f("abc", "abc"))
	assert.Equal(t, 0, f("abc", "def"))
}

func TestLongestCommonSubsequence(t *testing.T) {
	testLongestCommonSubsequence(t, solve1250.LongestCommonSubsequence)
	testLongestCommonSubsequence(t, solve1250.LongestCommonSubsequence2)
	testLongestCommonSubsequence(t, solve1250.LongestCommonSubsequence3)
}
