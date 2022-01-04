package solve0753_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0753"
)

func TestOpenLock(t *testing.T) {
	deadends := helper.ArrFromJSONString(`["0201","0101","0102","1212","2002"]`)
	target := "0202"
	expected := 6
	assert.Equal(t, expected, solve0753.OpenLock(deadends, target))
}
