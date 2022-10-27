package solve0349

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
)

func TestIntersection(t *testing.T) {
	assert.Equal(t, helper.ArrFromJSON("[2]"), Intersection(helper.ArrFromJSON("[1,2,2,1]"), helper.ArrFromJSON("[2,2]")))
}
