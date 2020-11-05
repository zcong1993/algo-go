package solve0753

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenLock(t *testing.T) {
	res := OpenLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	assert.Equal(t, 6, res)
}
