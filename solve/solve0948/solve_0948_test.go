package solve0948_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0948"
)

func TestSortArray(t *testing.T) {
	inputs := helper.ArrFromJSON(`[5,1,1,2,0,0]`)
	expected := helper.ArrFromJSON(`[0,0,1,1,2,5]`)
	assert.Equal(t, expected, solve0948.SortArray(inputs))
}
