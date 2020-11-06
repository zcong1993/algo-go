package solve1458

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func TestSortByBits(t *testing.T) {
	res := SortByBits(helper.ArrFromJSON("[1024,512,256,128,64,32,16,8,4,2,1]"))
	assert.Equal(t, helper.ArrFromJSON("[1,2,4,8,16,32,64,128,256,512,1024]"), res)
	assert.Equal(t, helper.ArrFromJSON("[0,1,2,4,8,3,5,6,7]"), SortByBits(helper.ArrFromJSON("[0,1,2,3,4,5,6,7,8]")))
}
