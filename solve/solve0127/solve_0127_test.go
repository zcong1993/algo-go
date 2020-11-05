package solve0127

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	assert.Equal(t, 5, LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	assert.Equal(t, 0, LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
